package plot

import (
	"fmt"
	"os"
)

const block = "▉"

// NewConsoleCanvas create canvas.
func NewConsoleCanvas(xSize, ySize int) Canvas {
	cubes := make([][]string, ySize)
	for i := range cubes {
		cubes[i] = make([]string, xSize)
	}

	return &consoleCanvas{
		xSize: xSize,
		ySize: ySize,

		cubes: cubes,
	}
}

// consoleCanvas impl Canvas interface, draw to console.
type consoleCanvas struct {
	xSize int
	ySize int

	cubes [][]string
}

// Clear Canvas.
func (c *consoleCanvas) Clear() {
	for i := 0; i < c.xSize; i++ {
		for j := 0; j < c.ySize; j++ {
			c.cubes[i][j] = "+"
		}
	}
}

// SetCubes to canvas.
func (c *consoleCanvas) SetCubes(cubes []*Cube) {
	for _, cube := range cubes {
		c.cubes[cube.Position.X][cube.Position.Y] = cube.Color.Sprint(block)
	}
}

// Flush cubes to screen.
func (c *consoleCanvas) Flush() {
	clearConsole()

	rawScreen := ""
	for i := 0; i < c.ySize; i++ {
		for j := 0; j < c.xSize; j++ {
			rawScreen += c.cubes[j][i]
		}
		rawScreen += "\n"
	}

	fmt.Fprint(os.Stdout, rawScreen)
}

func clearConsole() {
	fmt.Fprint(os.Stdout, "\x1bc")
}
