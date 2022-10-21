package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	noopGameScene := game.GameScene{}

	game := game.NewGame()
	game.AddScene("EMPTY", &noopGameScene)
	game.SetScene("EMPTY")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
