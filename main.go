package main

import (
	"image"
	"time"

	"github.com/fatih/color"
	"github.com/wang-sy/tetris/plot"
)

func main() {
	mainLoop()
}

func mainLoop() {
	canvas := plot.NewConsoleCanvas(20, 20)

	cubes := []*plot.Cube{
		{
			Position: image.Point{1, 1},
			Color:    color.New(color.FgBlue),
		},
	}

	for {
		time.Sleep(time.Millisecond * 100)

		canvas.Clear()
		canvas.SetCubes(cubes)
		canvas.Flush()

		setCube(cubes)
	}
}

func setCube(cubes []*plot.Cube) {
	cubes[0].Position.X += 1
	cubes[0].Position.X %= 20
}
