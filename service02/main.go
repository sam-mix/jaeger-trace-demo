package main

import (
	"context"
	"io"
	"log"
	"net"
	proto "trace-demo/cmd/protos"
	"trace-demo/tools"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	olog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	PORT = ":8082"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
)

type Server struct {
	proto.UnimplementedRpcServerServer
}

func (s *Server) Ping(ctx context.Context, all *proto.PingReq) (*proto.PingRes, error) {
	var res = new(proto.PingRes)
	span := ctx.Value("span").(opentracing.Span)
	span.LogFields(olog.String("pos", "rpcService:Ping"))
	res.Res = "pong"
	return res, nil
}

func main() {
	tracer, closer = tools.InitJaegerClient("trace-demo")
	defer closer.Close()

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := grpc.UnaryInterceptor(OpentracingServerInterceptor())
	s := grpc.NewServer(opts) //起一个服务
	proto.RegisterRpcServerServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func OpentracingServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(opentracing.TextMap, tools.MDReaderWriter{MD: md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err %v", err)
		}
		serverSpan := tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(spanContext),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
			ext.SpanKindRPCServer,
		)
		defer serverSpan.Finish()
		k := "span"
		ctx = context.WithValue(ctx, k, serverSpan)
		return handler(ctx, req)
	}
}
