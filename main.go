package main

import (
	"fmt"
	"syscall/js"
)

var mazeDisplay js.Value

func main() {
	mazeDisplay = js.Global().Get("document").Call("getElementById", "maze-display")
	mazeDisplay.Set("innerText", "This is where the maze will be displayed. If you don't see me it's not working!")
}
