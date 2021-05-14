.PHONY:proto
proto:
	@cd cmd/protos && \
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	ping.proto

.PHONY:build
build:
	@cd gateway && go build .  
	@cd service01 && go build . 	
	@cd service02 && go build . 	

.PHONY:clean
clean:
	@rm service01/service01
	@rm service02/service02
	@rm gateway/gateway

