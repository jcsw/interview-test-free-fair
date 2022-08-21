# Go parameters
GO_CMD=go
GO_GET=$(GO_CMD) get
GO_RUN=$(GO_CMD) run
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_VET=$(GO_CMD) vet
GO_COVER=$(GO_CMD) tool cover
GO_FMT=gofmt

BUILD_DIRECTORY=./build
BINARY_NAME=interview-test-free-fair-api
APP_INIT=./cmd/.

SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKG=$(shell go list ./... | grep -v /vendor/)

all: clean fmt vet build test 

build:
	mkdir -p $(BUILD_DIRECTORY)
	$(GO_BUILD) -o $(BUILD_DIRECTORY)/$(BINARY_NAME) -v $(APP_INIT)

test:
	$(GO_TEST) -v -short ./... -covermode=count -coverprofile=$(BUILD_DIRECTORY)/cover.out
	$(GO_COVER) -html=$(BUILD_DIRECTORY)/cover.out -o $(BUILD_DIRECTORY)/coverage.html
	xdg-open $(BUILD_DIRECTORY)/coverage.html &

itest:
	$(GO_TEST) -v ./... -covermode=count -coverprofile=$(BUILD_DIRECTORY)/icover.out
	$(GO_COVER) -html=$(BUILD_DIRECTORY)/icover.out -o $(BUILD_DIRECTORY)/icoverage.html
	xdg-open $(BUILD_DIRECTORY)/icoverage.html &

vet:
	$(GO_VET) $(PKG)

fmt:
	$(GO_FMT) -w -d $(SRC)

clean:
	$(GO_CLEAN)
	rm -rf $(BUILD_DIRECTORY)

run:
	$(GO_RUN) $(APP_INIT) -env=dev

deps-start:
	docker-compose -f docker/mariadb/docker-compose.yml up --remove-orphans

deps-stop:
	docker-compose -f docker/mariadb/docker-compose.yml stop

deps-rm:
	docker-compose -f docker/mariadb/docker-compose.yml rm
	rm -rf ~/.free_fair-mariadb:/var/lib/mysql

install:
	$(GO_GET) -d golang.org/x/tools/cmd/cover
	$(GO_GET) -d golang.org/x/lint
	$(GO_CMD) mod tidy
