build-protos:
	protoc -I protos protos/agent.proto --go_out=plugins=grpc:${GOPATH}/src

build-cli:
	make build-protos
	go build ./cmd/agent-cli

build-rpc:
	make build-protos
	go build ./cmd/agent-rpc

run-rpc:
	go run ./cmd/agent-rpc/main.go
