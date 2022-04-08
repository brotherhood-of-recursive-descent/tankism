SHELL := /bin/bash
	

run:
	go run app/tankism/main.go

test:
	go test ./...

build:
	go build -o dist/tankism app/tankism/main.go

dist: clean test build


clean:
	rm -rf dist/*
