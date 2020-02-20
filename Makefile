SHELL := bash
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
.PHONY: build clean format serve test vet

all_packages=
test_packages=

build: .make.main

clean:
	-rm main.wasm

format:
	gofmt -w main.go $(all_packages)

.make.main: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm
	touch .make.main

serve:
	goserve

test: format vet
	go test $(test_packages)

vet:
	go vet main.go $(all_packages)
