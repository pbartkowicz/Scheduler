groups=./data/groups.xlsx
students=./data/students
priority=./data/priority_students.xlsx
result=./data/result

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go -groups=$(groups) -students=$(students) -priority=$(priority) -result=$(result)

go-test:
	go test ./... -cover

clean:
	rm -rf bin

start:
	./bin/main -groups=$(groups) -students=$(students) -priority=$(priority) -result=$(result)

all: clean build go-test start