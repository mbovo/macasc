VERSION := $(shell git describe --tags | cut -d "-" -f 1)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
#GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
INSTALL_DIR := /usr/local/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
GOVENDOR=$(GOCMD) mod vendor
GOMOD=$(GOCMD) mod download
BIN_NAME=yacasc

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

#MAKEFLAGS += --silent
PHONY: clean deps build

all: test build
build: 
	@-mkdir -p $(GOBIN)
	$(GOBUILD) -o $(GOBIN)/$(BIN_NAME) $(LDFLAGS) -v cmd/yacasc.go

test:
	$(GOTEST) -v ./... -cover -coverprofile=c.out
	${GOCOVER} -html=c.out -o coverage.html
	@rm -rf c.out

clean:
	$(GOCLEAN)
	@-rm -rf $(GOBIN)
	@-rm -rf coverage.html

deps:
	$(GOMOD)

docker:
	docker build . -t yacasc:latest
install:
	cp $(GOBIN)/$(BIN_NAME) $(INSTALL_DIR)/$(BIN_NAME)