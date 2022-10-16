SHELL := /bin/bash
	

run:
	go run app/tankism/main.go

test:
	go test ./...

build:
	go build -o dist/tankism app/tankism/main.go

dist: clean test build

demo-lighting: 
	go run app/lighting/main.go

demo-animation: 
	go run app/animation/main.go

demo-loading: 
	go run app/loading/main.go

clean:
	rm -rf dist/*
