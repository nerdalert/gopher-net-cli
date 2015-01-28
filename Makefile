all: build test

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
		go test -covermode=count -coverprofile=main.cover.out