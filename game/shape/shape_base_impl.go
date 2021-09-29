package shape

import (
	"image"

	"github.com/fatih/color"

	"github.com/wang-sy/tetris/plot"
)

// baseShape impl shape related operations.
type baseShape struct {
	center image.Point
	color  *color.Color

	state       int
	relativePos [][]image.Point
}

func newBaseShapeFromRelativePos(center image.Point, color *color.Color, relativePos [][]image.Point) *baseShape {
	return &baseShape{
		center:      center,
		color:       color,
		state:       0,
		relativePos: relativePos,
	}
}

// ToCubes return nothing.
func (s *baseShape) ToCubes() []*plot.Cube {
	cubes := make([]*plot.Cube, len(s.relativePos[s.state]))

	for i, pos := range s.relativePos[s.state] {
		cubes[i] = &plot.Cube{
			Position: pos.Add(s.center),
			Color:    s.color,
		}
	}

	return cubes
}

// MoveDown return Shape after MoveDown.
func (s *baseShape) MoveDown() Shape {
	return &baseShape{
		center: image.Point{
			X: s.center.X,
			Y: s.center.Y + 1,
		},
		color:       s.color,
		state:       s.state,
		relativePos: s.relativePos,
	}
}

// MoveLeft return Shape after MoveLeft.
func (s *baseShape) MoveLeft() Shape {
	return &baseShape{
		center: image.Point{
			X: s.center.X - 1,
			Y: s.center.Y,
		},
		color:       s.color,
		state:       s.state,
		relativePos: s.relativePos,
	}
}

// MoveRight return Shape after MoveRight.
func (s *baseShape) MoveRight() Shape {
	return &baseShape{
		center: image.Point{
			X: s.center.X + 1,
			Y: s.center.Y,
		},
		color:       s.color,
		state:       s.state,
		relativePos: s.relativePos,
	}
}

// Roll return Shape after Roll.
func (s *baseShape) Roll() Shape {
	return &baseShape{
		center: image.Point{
			X: s.center.X,
			Y: s.center.Y,
		},
		color:       s.color,
		state:       (s.state + 1) % len(s.relativePos),
		relativePos: s.relativePos,
	}
}
