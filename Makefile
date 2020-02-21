SHELL := bash
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
.PHONY: clean format test vet

go_environment=GOOS=js GOARCH=wasm
all_packages=./cell ./grid
test_packages=./grid

build: vet format main.wasm

clean:
	-rm main.wasm

format:
	gofmt -w main.go $(all_packages)

main.wasm: main.go cell/cell.go grid/grid.go
	$(go_environment) go build -o main.wasm

serve: main.wasm
	goserve

test: vet format
	go test $(test_packages)

vet:
	$(go_environment) go vet $(all_packages)
	$(go_environment) go vet main.go
