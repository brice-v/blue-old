ifeq ($(OS),Windows_NT)
	EXE_NAME=blue.exe
else
	EXE_NAME=blue
endif

all: build test

build:
	go build -o ${EXE_NAME} .

test:
	go test ./...

testv:
	go test -v ./...

run: build
	./${EXE_NAME}

clean:
	go mod vendor
	go mod tidy
	go clean
