.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	go build -ldflags "-X github.com/devicechain-io/dcctl/cmd.gitCommit=$(GIT_COMMIT)" -o build/dcctl .
