all: build

clean:
	rm -rf bin/*

dependencies:
	go install

build: dependencies
	go build -o ./bin/hello

ci: dependencies test

build-mocks:
	@go install github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=db/queries.go -destination=mock/queries.go -package=mock

apidoc:
	@~/go/bin/apidoc -m ./main.go -o docs

test: build-mocks
	go test . -v

