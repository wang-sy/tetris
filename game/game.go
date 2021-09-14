package game

import (
	"fmt"
	"image"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/wang-sy/tetris/game/shape"
	"github.com/wang-sy/tetris/plot"
)

// Game describe resources needed for the game.
type Game struct {
	staticCubes []*plot.Cube
	activeShape shape.Shape

	xSize int
	ySize int

	cnt int
}

func New(xSize, ySize int) *Game {
	return &Game{
		staticCubes: make([]*plot.Cube, 0),
		activeShape: shape.NewLine(image.Point{10, 0}, color.New(color.FgBlue)),
		xSize:       xSize,
		ySize:       ySize,
		cnt:         0,
	}
}

func (g *Game) ProcessEvents(events []keyboard.KeyEvent) {
	for _, event := range events {
		activeShapeAfterEvent := g.posAfterEvent(event)
		if !g.validShapePos(activeShapeAfterEvent) {
			continue
		}

		g.activeShape = activeShapeAfterEvent
	}
}

func (g *Game) GetCubes() []*plot.Cube {
	cubes := make([]*plot.Cube, len(g.staticCubes))
	copy(cubes, g.staticCubes)

	cubes = append(cubes, g.activeShape.ToCubes()...)

	for _, cube := range cubes {
		fmt.Println(cube.Position, cube.Color.Sprint("a"))
	}

	return cubes
}

func (g *Game) validShapePos(s shape.Shape) bool {
	for _, cube := range s.ToCubes() {
		if cube.Position.X < 0 || cube.Position.X >= g.xSize {
			return false
		}

		if cube.Position.Y < 0 || cube.Position.Y >= g.ySize {
			return false
		}

		for _, staticCube := range g.staticCubes {
			if cube.Position.Eq(staticCube.Position) {
				return false
			}
		}
	}

	return true
}

func (g *Game) posAfterEvent(event keyboard.KeyEvent) shape.Shape {
	switch event.Key {
	case keyboard.KeyArrowDown:
		return g.activeShape.MoveDown()
	case keyboard.KeyArrowUp:
		return g.activeShape.Roll()
	case keyboard.KeyArrowLeft:
		return g.activeShape.MoveLeft()
	case keyboard.KeyArrowRight:
		return g.activeShape.MoveRight()
	default:
		return g.activeShape
	}
}
