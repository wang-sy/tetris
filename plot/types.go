package plot

import (
	"image"

	"github.com/fatih/color"
)

// Cube 描述一个方块.
type Cube struct {
	Position image.Point
	Color    *color.Color
}

// Canvas 描述画布.
type Canvas interface {
	Clear()
	SetCubes(cubes []*Cube)
	Flush()
}
