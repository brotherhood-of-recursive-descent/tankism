package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	emptyScene := &game.Scene{}

	client := game.NewGame()
	client.AddScene("EMPTY", emptyScene)

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
