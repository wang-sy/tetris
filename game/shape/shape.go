package shape

import (
	"image"

	"github.com/fatih/color"
	"github.com/wang-sy/tetris/plot"
)

// Shape describe a set of cubes.
type Shape interface {
	ToCubes() []*plot.Cube
	// MoveDown return Shape after MoveDown.
	MoveDown() Shape
	// MoveLeft return Shape after MoveLeft.
	MoveLeft() Shape
	// MoveRight return Shape after MoveRight.
	MoveRight() Shape
	// Roll return Shape after Roll.
	Roll() Shape
}

var lineRelativePos = [][]image.Point{
	{{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
	{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
}

// NewLineShape by center and color.
func NewLineShape(center image.Point, color *color.Color) Shape {
	return newBaseShapeFromRelativePos(center, color, lineRelativePos)
}
