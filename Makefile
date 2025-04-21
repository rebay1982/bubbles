.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet:
	go vet ./...

build: vet
	go build ./...

test: build
	go test -v ./...

run: build
	./hflogger
