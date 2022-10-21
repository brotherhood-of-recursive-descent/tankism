package main

import (
	_ "embed"
	"log"

	"github.com/co0p/tankism/game"

	"github.com/co0p/tankism/app/tankism/exit"
	"github.com/co0p/tankism/app/tankism/menu"
	"github.com/co0p/tankism/app/tankism/singleplayer"
	"github.com/co0p/tankism/app/tankism/start"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// construct game stuff
	game := game.NewGame()

	menuScene := menu.NewMenuScene(game)
	startScene := start.NewStartScreen(game)
	exitScene := exit.NewExitScene(game)
	singlePlayerScene := singleplayer.NewSinglePlayerScene(game)

	game.AddScene("MENU", menuScene)
	game.AddScene("START", startScene)
	game.AddScene("EXIT", exitScene)
	game.AddScene("SINGLEPLAYER", singlePlayerScene)

	game.SetScene("START")

	// setup
	ebiten.SetFullscreen(true)

	// start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
