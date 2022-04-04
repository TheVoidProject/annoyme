BINARY_NAME=annoyme
ANNOYME_LOCAL_DIR=${HOME}/.local/share/annoyme

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
	rm annoyme

install:
	mkdir -p ${ANNOYME_LOCAL_DIR}
	cp ./assets/icon.png ${ANNOYME_LOCAL_DIR}/icon.png

demo:
	./annoyme --daemon stop
	go build
	./annoyme --daemon start
	./annoyme new

all: dep build run clean