
BIN_DIR ?= ./bin
VERSION ?= $(shell git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty)
GO_LDFLAGS = -tags 'netgo osusergo static_build' -ldflags "-X github.com/edgefarm/edgefarm.network/cmd/credsmanager/cmd.version=$(VERSION)"
GO_ARCH = amd64

all: check test build

check:
	go mod vendor

build:
	GOOS=linux GOARCH=${GO_ARCH} go build ${GO_LDFLAGS} -o ${BIN_DIR}/credsmanager cmd/credsmanager/main.go

proto:
	go get -d google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I=pkg/apis/config/v1alpha1/ --go_out=pkg/apis/config/v1alpha1/ --go-grpc_out=pkg/apis/config/v1alpha1/ --go-grpc_opt=paths=source_relative pkg/apis/config/v1alpha1/config.proto

test:
	go test ./...

clean:
	rm -rf ${BIN_DIR}/credsmanager

.PHONY: check test clean build
