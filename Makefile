SHELL := bash
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
.PHONY: test format vet

all_packages=
test_packages=

test: format vet
	go test $(test_packages)

format:
	gofmt -w main.go $(all_packages)

vet:
	go vet main.go $(all_packages)
