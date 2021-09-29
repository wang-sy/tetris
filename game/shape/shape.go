package shape

import (
	"image"
	"log"
	"math/rand"

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

// NewRandomShape create a random shape.
func NewRandomShape(center image.Point, color *color.Color) Shape {
	shapeNameIndex := rand.Int() % len(shapeNames)

	return NewShapeFromName(shapeNames[shapeNameIndex], center, color)
}

// NewShapeFromName create shape from it's name.
func NewShapeFromName(name string, center image.Point, color *color.Color) Shape {
	shapeRelativePos, ok := shapeRelativePosMap[name]
	if !ok {
		log.Fatal("unsupport name:", name)
	}

	return newBaseShapeFromRelativePos(center, color, shapeRelativePos)
}

var shapeNames = []string{lineName, lShapeName, tShapeName, squareName}

var shapeRelativePosMap = map[string][][]image.Point{
	lineName:   lineRelativePos,
	lShapeName: lShapeRelativePos,
	squareName: squareRelativePos,
	tShapeName: tShapeRelativePos,
}

const (
	lineName   = "line"
	lShapeName = "lShape"
	tShapeName = "tShape"
	squareName = "square"
)

var (
	lineRelativePos = [][]image.Point{
		{{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
	}
	lShapeRelativePos = [][]image.Point{
		{{1, 0}, {0, 0}, {0, 1}, {0, 2}},
		{{0, -1}, {0, 0}, {1, 0}, {2, 0}},
		{{-1, 0}, {0, 0}, {0, -1}, {0, -2}},
		{{-2, 0}, {-1, 0}, {0, 0}, {0, 1}},
	}
	tShapeRelativePos = [][]image.Point{
		{{-1, 0}, {0, 0}, {1, 0}, {0, 1}},
		{{0, 0}, {0, 1}, {0, -1}, {1, 0}},
		{{-1, 0}, {0, 0}, {1, 0}, {0, -1}},
		{{-1, 0}, {0, 0}, {0, 1}, {0, -1}},
	}
	squareRelativePos = [][]image.Point{
		{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
	}
)
