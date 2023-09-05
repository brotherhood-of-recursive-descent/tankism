package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/game/state"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/test"
	"github.com/hajimehoshi/ebiten/v2"
)

type LoadingDemo struct {
	game.GameScene
}

func NewLoadingDemo() *LoadingDemo {
	demo := LoadingDemo{}

	state, _ := state.New(test.GameState_Valid)
	demo.EntityManager = *ecs.NewEntityManager(state.Entities)

	demo.Systems = append(demo.Systems,
		&systems.SpriteRenderer{EntityManager: &demo.EntityManager},
	)
	return nil
}

func main() {
	demo := LoadingDemo{}

	game := game.NewGame()
	game.AddScene("Loading", &demo)
	game.SetScene("Loading")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
