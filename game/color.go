package game

import (
	"math/rand"

	"github.com/fatih/color"
)

func newRandomColor() *color.Color {
	colorIndex := rand.Int() % len(colors)

	return color.New(colors[colorIndex])
}

var colors = []color.Attribute{
	color.FgRed,
	color.FgGreen,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,
}
