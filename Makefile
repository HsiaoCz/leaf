run: build
	@./bin/leaf

build:
	@go build -o bin/leaf main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

all: run

help:
	@echo "run: run project"
	@echo "build: go build -o main.go bin/leaf"
	@echo "if input make with nothing will use run"
	@echo "tidy: go mod tidy"

.PHONY:run build test tidy help all
