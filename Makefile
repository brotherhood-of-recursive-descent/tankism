SHELL := /bin/bash


run:
	go run app/tankism/main.go

test:
	go test ./...

build:
	go build -o tankism app/tankism/main.go

clean:
	rm -rf tankism
