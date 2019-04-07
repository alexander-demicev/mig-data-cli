.PHONY: build build-alpine clean test help default

BIN_NAME=ocp-mig-test-data-cli

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "ademicev/ocp-mig-test-data-cli"

default: test

help:
	@echo 'Management commands for ocp-mig-test-data-cli:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/alexander-demichev/ocp-mig-test-data-cli/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/alexander-demichev/ocp-mig-test-data-cli/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

