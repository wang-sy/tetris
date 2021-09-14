package shape

import (
	"image"

	"github.com/fatih/color"
	"github.com/wang-sy/tetris/plot"
)

var lineRelativePos = [][]image.Point{
	{{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
	{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
}

type lineShape struct {
	// center -*--.
	center image.Point
	color  *color.Color

	state int
}

// NewLine shape.
func NewLine(center image.Point, color *color.Color) Shape {
	return &lineShape{
		center: center,
		color:  color,
		state:  0,
	}
}

// ToCubes.
func (s *lineShape) ToCubes() []*plot.Cube {
	cubes := make([]*plot.Cube, len(lineRelativePos[s.state]))

	for i, pos := range lineRelativePos[s.state] {
		cubes[i] = &plot.Cube{
			Position: pos.Add(s.center),
			Color:    s.color,
		}
	}

	return cubes
}

// MoveDown return Shape after MoveDown.
func (s *lineShape) MoveDown() Shape {
	return &lineShape{
		center: image.Point{
			X: s.center.X,
			Y: s.center.Y + 1,
		},
		color: s.color,
		state: s.state,
	}
}

// MoveLeft return Shape after MoveLeft.
func (s *lineShape) MoveLeft() Shape {
	return &lineShape{
		center: image.Point{
			X: s.center.X - 1,
			Y: s.center.Y,
		},
		color: s.color,
		state: s.state,
	}
}

// MoveRight return Shape after MoveRight.
func (s *lineShape) MoveRight() Shape {
	return &lineShape{
		center: image.Point{
			X: s.center.X + 1,
			Y: s.center.Y,
		},
		color: s.color,
		state: s.state,
	}
}

// Roll return Shape after Roll.
func (s *lineShape) Roll() Shape {
	return &lineShape{
		center: image.Point{
			X: s.center.X,
			Y: s.center.Y,
		},
		color: s.color,
		state: (s.state + 1) % len(lineRelativePos),
	}
}
