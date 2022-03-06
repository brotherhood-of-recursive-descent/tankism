SHELL := /bin/bash


run:
	go run cmd/tankism/main.go

test:
	go test ./...

build:
	go build -o tankism cmd/tankism/main.go

clean:
	rm -rf tankism
