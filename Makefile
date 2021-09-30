all: build

BIN = $(CURDIR)/bin

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
	@GOBIN=$(GOBIN) go install github.com/golang/mock/mockgen@v1.6.0
	@$(GOBIN)/mockgen -source=db/queries.go -destination=mock/queries.go -package=mock

apidoc:
	@$(GOBIN)/apidoc -m ./main.go -o docs

check:
	@golangci-lint run

test: unit-tests integration-tests

unit-tests: build-mocks
	@CGO_ENABLED="0" go test ./app -v

integration-tests:
	@go test -cpu 1 -count=1 -v ./integration_tests

