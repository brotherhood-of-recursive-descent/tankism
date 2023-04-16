package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type RotationGame struct {
	game.GameScene
}

func (s *RotationGame) Init() error {

	s.Systems = append(s.Systems,
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.RotationSystem{EntityManager: &s.EntityManager},
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
	)

	fps := s.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	withoutRotation := s.EntityManager.NewEntity()
	game.NewTankWithPosition(withoutRotation, 200, 200)

	withRotation := s.EntityManager.NewEntity()
	game.NewTankWithPosition(withRotation, 400, 200)
	sprite := withRotation.GetComponent(components.SpriteType).(*components.Sprite)
	w, h := sprite.Image.Size()

	rotation := components.Rotation{
		Point: vector.Vec2d{
			X: float64(w / 2),
			Y: float64(h / 2),
		},
	}
	withRotation.AddComponent(&rotation)

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
