package main

import (
	"github.com/samwson/go-maze/grid"
	"image"
	"syscall/js"
)

var (
	mazeDisplay js.Value
	maze        grid.Grid
)

func main() {
	extent := image.Point{X: 10, Y: 10}
	maze := grid.New(extent)

	mazeDisplay = js.Global().Get("document").Call("getElementById", "maze-display")
	mazeDisplay.Set("innerText", maze.String())
}
