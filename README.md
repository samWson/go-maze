# Go Maze

An exercise in maze generating algorithms using WebAssembly and Go.

## Description

This is a WebAssembly app with Go doing all the maze generation. A minimal HTML
and JavaScript front end displays the generated mazes.

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
our work hackdays the copyrights belong to Flux Federation Ltd. The text of the license
is in the `LICENSE` file. Excluded from these conditions is some derived source
code listed below.

The files `index.html` and `wasm_exec.js` are derived from files in the `$GOROOT/misc/wasm`
directory of the Go `1.13.4` distribution. They retain their copyrights to The Go
Authors and are under BSD-style license. The text of the license is in the `GO_AUTHORS_LICENSE`
file. The files have the copyrights listed in their headers.
