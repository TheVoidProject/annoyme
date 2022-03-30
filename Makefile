BINARY_NAME=annoyme

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	go build -o ./bin/${BINARY_NAME}

run:
	./bin/${BINARY_NAME}

clean:
	rm -rf bin/*
	rm -f annoyme-sqlite.db

dep:
	go mod vendor

lint:
	golangci-lint run --enable-all

install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

all:
	dep build run clean
