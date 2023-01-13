package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionDemo struct {
	game.GameScene
}

func (demo *CollisionDemo) Init() error {

	demo.Systems = append(demo.Systems,
		&systems.SpriteRenderer{EntityManager: &demo.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &demo.EntityManager},
		&systems.TextRenderer{EntityManager: &demo.EntityManager},
		&systems.Controller{EntityManager: &demo.EntityManager},
	)

	fps := demo.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	// add some items
	barrel := demo.EntityManager.NewEntity()
	game.NewDrum(barrel, 300, 300)

	crate := demo.EntityManager.NewEntity()
	game.NewCrate(crate, 100, 300)

	tank := demo.EntityManager.NewEntity()
	game.NewTank(tank)

	return nil
}

func main() {

	demo := CollisionDemo{}
	game := game.NewGame()
	game.AddScene("CollisionDemo", &demo)
	game.SetScene("CollisionDemo")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
