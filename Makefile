BIN_PATH ?= ./bin/service
GENERATED_FILE_PATTERNS = *_easyjson.go *_gen.go *_get_test.go

all: dep fmt gen build

fmt:
	go fmt ./...

gen:
	go generate ./...

dep:
	go get -u && go mod tidy && go mod vendor && go mod verify

build: mkdir
	go build -mod vendor -o $(BIN_PATH) .

image:
	./image.sh build

test:
	go test -mod vendor ./...

clean:
	rm -rf ./bin

mkdir:
	mkdir -p ./bin
