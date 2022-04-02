BINARY_NAME=annoyme

.PHONY: all

dep:
	go mod download
	go mod vendor

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	go build -o ./bin/${BINARY_NAME}

run:
	./bin/${BINARY_NAME}

clean:
	rm -rf bin/*
	rm -rf logs/*

all: dep build run clean