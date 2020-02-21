package cell

import "image"

const (
	north = "north"
	east  = "east"
	south = "south"
	west  = "west"
)

type Cell struct {
	position image.Point
	Links    []string
}
