all: build test

install:
		go install golang.org/x/tools/cmd/vet

build:
		go build -v ./...
		go fmt ./...
		go vet ./...
		go test -x -v ./...
		go test -i -race ./...
		go test -v -race ./...
		go get -d -v ./... && go build

test:
		go test -covermode=count -coverprofile=main.cover.out -test.short

test-all:
		make install
		make build
		go test -covermode=count -coverprofile=main.cover.out
