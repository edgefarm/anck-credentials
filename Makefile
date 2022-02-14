
BIN_DIR ?= ./bin
VERSION ?= $(shell git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty)
GO_LDFLAGS = -tags 'netgo osusergo static_build' -ldflags "-X github.com/edgefarm/anck-credentials/cmd/anck-credentials/cmd.version=$(VERSION)"
GO_ARCH = amd64

all: tidy test build ## default target: tidy, test, build

tidy: ## ensure that go dependencies are up to date
	go mod vendor

build: ## build anck-credentials
	GOOS=linux GOARCH=${GO_ARCH} go build ${GO_LDFLAGS} -o ${BIN_DIR}/anck-credentials cmd/anck-credentials/main.go

proto: ## generate proto files
	go get -d google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I=pkg/apis/config/v1alpha1/ --go_out=pkg/apis/config/v1alpha1/ --go-grpc_out=pkg/apis/config/v1alpha1/ --go-grpc_opt=paths=source_relative pkg/apis/config/v1alpha1/config.proto

test: ## run go tests
	go test ./...

clean: ## cleans all
	rm -rf ${BIN_DIR}/anck-credentials

.PHONY: all tidiy build proto tes clean help

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m\t %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
