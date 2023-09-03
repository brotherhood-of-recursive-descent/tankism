package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type RotationGame struct {
	game.GameScene
}

func (s *RotationGame) Init() error {

	s.Systems = append(s.Systems,
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.MotionControlSystem{EntityManager: &s.EntityManager},
		&systems.MotionSystem{EntityManager: &s.EntityManager},
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
	)

	fps := s.EntityManager.NewEntity()
	tank := s.EntityManager.NewEntity()
	game.NewTankWithPosition(tank, 200, 200)
	tank.AddComponent(components.NewMotionControl())

	game.FPSCounter(fps, 1024)

	return nil
}

func main() {

	game := game.NewGame()
	demo := RotationGame{}

	game.AddScene("RotationGame", &demo)
	game.SetScene("RotationGame")

	ebiten.SetFullscreen(false)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
