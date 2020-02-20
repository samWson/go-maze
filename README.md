# Go Maze

An exercise in maze generating algorithms using WebAssembly and Go.

## Usage

A Makefile has the tasks required for building and running the code.

```shell
# Build the *.wasm files
$ make build

# Run the server on localhost:8080
$ make serve
```

The make file uses the `goserve` command. This is a simple HTTP server based on
this code [here](https://play.golang.org/p/pZ1f5pICVbV). Alternatively the [`goexec`](https://github.com/shurcooL/goexec#goexec)
command can be installed and used to run a server instead:
```shell
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

## License

This repository is open source software under an MIT license. As it was made during
our work hackdays the copyrights belong to Flux Federation Ltd.
