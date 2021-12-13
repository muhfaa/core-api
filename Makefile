BINARY=engine

dependencies:
	go mod tidy

build: dependencies
	GOOS=linux GOARCH=amd64 go build -o majoo app/main.go

run:
	go run app/main.go


.PHONY: dependencies build run