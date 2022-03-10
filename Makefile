.PHONY: gen-proto
get-proto:
	protoc --go_out=./gen-client --go_opt=paths=source_relative \
    --go-grpc_out=./gen-client --go-grpc_opt=paths=source_relative \
	proto/api.proto
