BIN_PATH ?= ./bin/service

REPOSITORIES_ENABLED ?= true
QUEUES_ENABLED ?= true

LD_FLAGS := "-X config.repositoriesEnabled=$(REPOSITORIES_ENABLED) -X config.queuesEnabled=$(QUEUES_ENABLED)"



all: dep fmt gen build

fmt:
	go fmt ./...


gen:
	go generate ./...

	mockgen -source=repositories/usersRepository.go \
		-destination=repositories/mock/usersRepository.go
	mockgen -source=repositories/documentsInfoRepository.go \
		-destination=repositories/mock/documentsInfoRepository.go
	mockgen -source=repositories/documentsContentRepository.go \
		-destination=repositories/mock/documentsContentRepository.go

	mockgen -source=logic/usersService.go \
		-destination=logic/mock/usersService.go
	mockgen -source=logic/documentsInfoService.go \
    	-destination=logic/mock/documentsInfoService.go
	mockgen -source=logic/documentsContentService.go \
		-destination=logic/mock/documentsContentService.go


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
