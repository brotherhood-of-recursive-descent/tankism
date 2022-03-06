package main

import (
	"github.com/co0p/tankism"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	// setup
	ebiten.SetFullscreen(true)

	// construct game stuff
	client := tankism.NewClient()

	// start the game
	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
