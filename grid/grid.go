package grid

import (
	"fmt"
	"github.com/samwson/go-maze/cell"
	"image"
	"strings"
)

const (
	corner                 = "+"
	horizontal             = "---"
	horizontalToNextCorner = "---+"
	vertical               = "|"
	empty                  = "   " // Three white spaces
	emptyToNextVertical    = "   |"
)

type Grid struct {
	extent image.Point
	Cells  []cell.Cell
}

func New(extent image.Point) *Grid {
	size := extent.X * extent.Y

	return &Grid{
		Cells:  make([]cell.Cell, size),
		extent: extent,
	}
}

func (g *Grid) String() string {
	var builder strings.Builder
	rowBoundary := fmt.Sprint(corner, strings.Repeat(horizontalToNextCorner, g.columns()), "\n")
	rowBody := fmt.Sprint(vertical, strings.Repeat(emptyToNextVertical, g.columns()), "\n")

	builder.WriteString(rowBoundary)

	for row := 0; row < g.rows(); row++ {
		builder.WriteString(rowBody)
		builder.WriteString(rowBoundary)
	}

	return strings.TrimSpace(builder.String())
}

func (g *Grid) columns() int {
	return g.extent.X
}

func (g *Grid) rows() int {
	return g.extent.Y
}
