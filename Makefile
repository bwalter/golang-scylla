all: build

BIN = $(CURDIR)/bin

ifndef GOPATH
GOPATH := $(HOME)/bin
endif

ifndef GOBIN
GOBIN := $(GOPATH)/bin
endif

clean:
	@rm -rf $(BIN)/*

dependencies:
	@GOBIN=$(GOBIN) go install

build: dependencies
	@mkdir -p $(BIN)
	@go build -o $(BIN)/hello

build-mocks:
	@rm -f ./mock/*
	@GOBIN=$(GOBIN) go install github.com/golang/mock/mockgen@v1.6.0
	@$(GOBIN)/mockgen -source=db/database.go -destination=mock/database.go -package=mock

apidoc:
	@GOBIN=$(GOBIN) go install github.com/spaceavocado/apidoc@v0.3.5
	@$(GOBIN)/apidoc -m ./main.go -o doc

check:
	@golangci-lint run

test: unit-tests integration-tests

unit-tests: build-mocks
	@CGO_ENABLED="0" go test ./app -v
	@CGO_ENABLED="0" go test ./routing -v

integration-tests:
	@go test -cpu 1 -count=1 -v ./integration_tests

