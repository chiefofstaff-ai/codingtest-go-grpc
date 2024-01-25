# Install the go code generator
.PHONY: install.grpc
grpc.install:
	brew install protobuf && brew install protoc-gen-go

.SILENT: grpc.gen
.PHONY: grpc.gen
grpc.gen:
	protoc cos.proto --proto_path=protos \
		--go_out=protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=protos \
		--go-grpc_opt=paths=source_relative

run:
	go run main.go
