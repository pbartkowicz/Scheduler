build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

go-test:
	go test ./... -cover

clean:
	rm -rf bin

start:
	./bin/main

all: clean build go-test start