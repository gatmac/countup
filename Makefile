.DEFAULT_GOAL := build
BINARY_NAME=countup

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet

build: vet
	go build -o bin/${BINARY_NAME} .
.PHONY:build

install: build
	@if [ -d ${GOPATH}/bin ]; then cp bin/${BINARY_NAME} ${GOPATH}/bin; fi
.PHONY:install

run: vet
	go run .
.PHONY:run

clean:
	rm bin/*
	go clean
.PHONY:clean

compile: vet
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64.exe .
.PHONY:compile