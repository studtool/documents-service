BIN_PATH ?= ./bin/service

SERVICE_NAME := documents-service
REPOSITORIES_ENABLED ?= true
QUEUES_ENABLED ?= true

LD_FLAGS := "-X config.Component=${SERVICE_NAME} -X config.repositoriesEnabled=$(REPOSITORIES_ENABLED) -X config.queuesEnabled=$(QUEUES_ENABLED)"

all: dep fmt gen build

fmt:
	go fmt ./...

gen:
	go generate ./...

dep:
	go get -u && go mod tidy && go mod vendor && go mod verify

build: mkdir
	go build -mod vendor -ldflags $(LD_FLAGS) -o $(BIN_PATH) .

image:
	./image.sh build

test:
	go test -mod vendor ./...

clean:
	rm -rf ./bin

mkdir:
	mkdir -p ./bin
