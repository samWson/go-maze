package grid

import (
	"image"
	"strings"
	"testing"
)

func TestUnlinkedGrid(t *testing.T) {
	extent := image.Point{X: 4, Y: 3}

	grid := New(extent)

	expected :=
		`+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+`

	actual := grid.String()

	if actual != strings.TrimSpace(expected) {
		t.Errorf("`grid.String()` does not match expected value. Expected= \n%s\nActual= \n%s\n", expected, actual)
	}
}
