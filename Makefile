
BIN_DIR ?= ./bin
VERSION ?= $(shell git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty)
GO_LDFLAGS = -tags 'netgo osusergo static_build'
GO_ARCH = amd64

all: check test build

build:
	GOOS=linux GOARCH=${GO_ARCH} go build -o ${BIN_DIR}/credsmanager cmd/credsmanager/main.go

proto:
	go get -d google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I=pkg/config/v1alpha1 --go_out=pkg/config/v1alpha1 --go-grpc_out=pkg/config/v1alpha1 --go-grpc_opt=paths=source_relative pkg/config/v1alpha1/api.proto

test:
	go test ./...

clean:
	rm -rf ${BIN_DIR}/credsmanager

.PHONY: check test clean

