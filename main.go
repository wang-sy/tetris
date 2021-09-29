package main

import (
	"log"
	"time"

	"github.com/wang-sy/tetris/game"
	"github.com/wang-sy/tetris/keyboard"
	"github.com/wang-sy/tetris/plot"
)

const (
	xSize = 10
	ySize = 10
)

func main() {
	events, err := keyboard.NewKeyboardEventBuffer()
	if err != nil {
		log.Fatal(err)
	}

	mainLoop(events)
}

func mainLoop(eventWatcher *keyboard.KeyboardEventBuffer) {
	gameController := game.New(xSize, ySize)
	canvas := plot.NewConsoleCanvas(xSize, ySize)

	go func() {
		for {
			time.Sleep(time.Millisecond * 400)
			gameController.MoveDown()
		}
	}()

	for {
		events := eventWatcher.ListAndClear()
		gameController.ProcessEvents(events)

		canvas.Clear()
		canvas.SetCubes(gameController.GetCubes())
		canvas.Flush()

		time.Sleep(time.Millisecond * 10)
	}
}
