BINARY 	      := efishery-cli-app
BUILD_DIR   	:= build
CGO_ENABLED   := 1
CGO_CFLAGS    := "-g -O2 -Wno-return-local-addr"

ifndef GOOS
  GOOS := $(shell go env GOHOSTOS)
endif

ifndef GOARCH
	GOARCH := $(shell go env GOHOSTARCH)
endif

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) CGO_CFLAGS=$(CGO_CFLAGS) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BUILD_DIR)/$(BINARY)
