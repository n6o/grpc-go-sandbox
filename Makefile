
.PHONY: gen-proto
gen-proto:
	protoc \
	-I . \
	-I ./vendor/github.com/envoyproxy/protoc-gen-validate \
	--go_out=./gen-client --go_opt=paths=source_relative \
	--go-grpc_out=./gen-client --go-grpc_opt=paths=source_relative \
	--validate_out=paths=source_relative,lang=go:./gen-client \
	proto/api.proto

.PHONY: gen-mock
gen-mock:
	mockgen -source ./gen-client/proto/api_grpc.pb.go -package gen_client -destination ./gen-client/proto/mock_api_grpc.pb.go
	mockgen -source ./gen-client/proto/api.pb.go -package gen_client -destination ./gen-client/proto/mock_api.pb.go
