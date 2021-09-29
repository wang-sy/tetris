package game

import (
	"image"
	"log"

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
	g := &Game{
		staticCubes: make([]*plot.Cube, 0),
		xSize:       xSize,
		ySize:       ySize,
		cnt:         0,
	}

	g.generateNewActiveShape()

	return g
}

func (g *Game) ProcessEvents(events []keyboard.KeyEvent) {
	for _, event := range events {
		activeShapeAfterEvent := g.posAfterEvent(event)
		if g.validShapePos(activeShapeAfterEvent) {
			g.activeShape = activeShapeAfterEvent

			continue
		}

		// not valid and event is moveDown, do post moveDown.
		if event.Key == keyboard.KeyArrowDown {
			g.mergeActiveShapeToStatic()
			g.generateNewActiveShape()
		}
	}
}

func (g *Game) GetCubes() []*plot.Cube {
	cubes := make([]*plot.Cube, len(g.staticCubes))
	copy(cubes, g.staticCubes)

	cubes = append(cubes, g.activeShape.ToCubes()...)

	return cubes
}

func (g *Game) MoveDown() {
	activeShapeAfterEvent := g.activeShape.MoveDown()
	if g.validShapePos(activeShapeAfterEvent) {
		g.activeShape = activeShapeAfterEvent

		return
	}

	g.mergeActiveShapeToStatic()
	g.generateNewActiveShape()
}

func (g *Game) mergeActiveShapeToStatic() {
	g.staticCubes = append(g.staticCubes, g.activeShape.ToCubes()...)

	yCount := make([]int, g.ySize)

	for _, cube := range g.staticCubes {
		yCount[cube.Position.Y]++
	}

	ySum := make([]int, g.ySize)

	pre := 0
	for i := g.ySize - 1; i >= 0; i-- {
		if yCount[i] == g.xSize {
			ySum[i] = 1
		} else {
			ySum[i] = 0
		}

		ySum[i] += pre
		pre = ySum[i]
	}

	newCubes := make([]*plot.Cube, 0)
	for _, cube := range g.staticCubes {
		if yCount[cube.Position.Y] == g.xSize {
			continue
		}

		newCubes = append(newCubes, &plot.Cube{
			Position: image.Pt(cube.Position.X, cube.Position.Y+ySum[cube.Position.Y]),
			Color:    cube.Color,
		})
	}

	g.staticCubes = newCubes
}

func (g *Game) generateNewActiveShape() {
	g.activeShape = shape.NewRandomShape(image.Point{g.xSize / 2, 0}, color.New(color.FgBlue))

	if !g.validShapePos(g.activeShape) {
		log.Fatal("game over")
	}
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
