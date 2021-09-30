all: build

clean:
	@rm -rf bin/*

dependencies:
	@go install

build: dependencies
	@go build -o ./bin/hello

build-mocks:
	@go install github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=db/queries.go -destination=mock/queries.go -package=mock

apidoc:
	@~/go/bin/apidoc -m ./main.go -o docs

check:
	@golangci-lint run

test: unit-tests integration-tests

unit-tests: build-mocks
	@go test ./app -v

integration-tests:
	@go test -cpu 1 -count=1 -v ./integration_tests

