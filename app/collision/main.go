package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
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
		&systems.MotionControlSystem{EntityManager: &demo.EntityManager},
		&systems.Shaker{EntityManager: &demo.EntityManager},
		&systems.CollisionDetection{EntityManager: &demo.EntityManager},
		&systems.CollisionResolution{EntityManager: &demo.EntityManager},
		&systems.DebugRenderer{EntityManager: &demo.EntityManager},
	)

	fps := demo.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	barrel := demo.EntityManager.NewEntity()
	game.NewDrum(barrel, 300, 300)

	crate := demo.EntityManager.NewEntity()
	game.NewCrate(crate, 100, 300)

	tank := demo.EntityManager.NewEntity()
	game.NewTank(tank)
	tank.AddComponent(&components.Debug{})

	tree := demo.EntityManager.NewEntity()
	game.NewTree(tree, 500, 300)
	tree.AddComponent(&components.Debug{})

	bullet := demo.EntityManager.NewEntity()
	game.NewBullet(bullet, 700, 300)

	bigTank := demo.EntityManager.NewEntity()
	game.NewBigTank(bigTank, 900, 300)

	oilSpill := demo.EntityManager.NewEntity()
	game.NewOilSpill(oilSpill, 300, 600)

	barricade := demo.EntityManager.NewEntity()
	game.NewBarricade(barricade, 600, 600)

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
