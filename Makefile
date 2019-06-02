BIN_PATH ?= ./bin/service

LOGS_EXPORT_ENABLED ?= true
REPOSITORIES_ENABLED ?= true
QUEUES_ENABLED ?= true

ORIGIN := github.com
APP_NAME := studtool

SERVICE_NAME := documents-service
SERVICE_VERSION := v0.0.1

APP_PACKAGE := $(ORIGIN)/$(APP_NAME)/$(SERVICE_NAME)

LD_FLAGS := -X $(APP_PACKAGE)/config.ComponentName=$(SERVICE_NAME) \
			-X $(APP_PACKAGE)/config.ComponentVersion=$(SERVICE_VERSION) \
			-X $(APP_PACKAGE)/config.logsExportEnabled=$(LOGS_EXPORT_ENABLED) \
			-X $(APP_PACKAGE)/config.repositoriesEnabled=$(REPOSITORIES_ENABLED) \
			-X $(APP_PACKAGE)/config.queuesEnabled=$(QUEUES_ENABLED)

IMAGE_TAG := $(APP_NAME)/$(SERVICE_NAME):$(SERVICE_VERSION)

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
	go build -mod vendor -ldflags '$(LD_FLAGS)' -o '$(BIN_PATH)' .

build_image:
	./image.sh build '$(IMAGE_TAG)'

push_image:
	./image.sh push '$(IMAGE_TAG)'

test:
	go test -mod vendor ./...

lint:
	./linter.sh run

clean:
	rm -rf ./bin

mkdir:
	mkdir -p ./bin
