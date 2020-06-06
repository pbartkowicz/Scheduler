build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go -gf=$(gf) -sd=$(sd) -psf=$(psf) -rd=$(rd)

go-test:
	go test ./... -cover

clean:
	rm -rf bin

start:
	./bin/main -gf=$(gf) -sd=$(sd) -psf=$(psf) -rd=$(rd)

all: clean build go-test start