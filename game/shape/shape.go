package shape

import "github.com/wang-sy/tetris/plot"

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
