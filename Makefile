# Meta info
NAME := go-fm
VERSION := v$(shell cat VERSION)
MAINTAINER := "Vicente Zepeda <chente.z.m@gmail.com"
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)

all: install_dependencies test

lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)

test: lint
	 @go test -v $(SOURCE_DIRS) -cover -bench . -race

install_dependencies:
	dep ensure

cover:
	@bash cover.sh

.PHONY: lint test install_dependencies cover