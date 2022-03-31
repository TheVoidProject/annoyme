# BINARY_NAME=annoyme

# .PHONY: all

# dep:
# 	go mod download
# 	go mod vendor

# install-linter:
# 	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# lint:
# 	golangci-lint run 
	
# #--enable-all


# build:
# 	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
# 	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
# 	go build -o ./bin/${BINARY_NAME}

# run:
# 	./bin/${BINARY_NAME}

# clean:
# 	rm -rf bin/*
# 	rm -f annoyme-sqlite.db

# all: dep install-linter lint build run clean