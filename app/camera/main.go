package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type CameraDemo struct {
	game.GameScene
	game *game.Game
}

func (s *CameraDemo) Init() error {

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.MovementSystem{EntityManager: &s.EntityManager},

		systems.NewCameraSystem(&s.EntityManager, s.game.ScreenWidth, s.game.ScreenHeight),
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
	)

	fps := s.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	tank := s.EntityManager.NewEntity()
	game.NewTankWithPosition(tank, 400, 400)
	tank.AddComponent(&components.Camera{})

	barrel := s.EntityManager.NewEntity()
	game.NewDrum(barrel, 100, 100)

	barrel2 := s.EntityManager.NewEntity()
	game.NewDrum(barrel2, 1000, 100)

	barrel3 := s.EntityManager.NewEntity()
	game.NewDrum(barrel3, 100, 1000)

	barrel4 := s.EntityManager.NewEntity()
	game.NewDrum(barrel4, 1000, 1000)

	return nil
}

func main() {

	game := game.NewGame()
	demo := CameraDemo{game: game}

	game.AddScene("Demo", &demo)
	game.SetScene("Demo")

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
