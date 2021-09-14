package main

import (
	"log"
	"time"

	"github.com/wang-sy/tetris/game"
	"github.com/wang-sy/tetris/keyboard"
	"github.com/wang-sy/tetris/plot"
)

func main() {
	events, err := keyboard.NewKeyboardEventBuffer()
	if err != nil {
		log.Fatal(err)
	}

	mainLoop(events)
}

func mainLoop(eventWatcher *keyboard.KeyboardEventBuffer) {
	gameController := game.New(20, 20)
	canvas := plot.NewConsoleCanvas(20, 20)

	for {
		events := eventWatcher.ListAndClear()
		gameController.ProcessEvents(events)

		canvas.Clear()
		canvas.SetCubes(gameController.GetCubes())
		canvas.Flush()

		time.Sleep(time.Millisecond * 100)
	}
}
