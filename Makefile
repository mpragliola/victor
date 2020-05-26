.PHONY: build

run:
	go run ./cmd build

build:
	go build -o bin/ cmd/victor.go