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
}

func (s *CameraDemo) Init() error {

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.MovementSystem{EntityManager: &s.EntityManager},
		&systems.CameraSystem{EntityManager: &s.EntityManager},
	)

	fps := s.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	tank := s.EntityManager.NewEntity()
	game.NewTankWithPosition(tank, 400, 400)
	tank.AddComponent(&components.Camera{X: 400, Y: 400})

	tileMap := s.EntityManager.NewEntity()
	game.NewMap(tileMap, game.Tilemap{}, 1024, 1024)

	return nil
}

func (s *CameraDemo) HandleInput() {
	x, y := ebiten.CursorPosition()
	entities := s.EntityManager.FindByComponents(components.CameraType)
	if len(entities) != 1 {
		log.Fatalf("expected to have camera in scene")
	}

	camera := entities[0].GetComponent(components.CameraType).(*components.Camera)
	camera.X = float64(x)
	camera.Y = float64(y)
}

func main() {

	game := game.NewGame()
	demo := CameraDemo{}

	game.AddScene("Demo", &demo)
	game.SetScene("Demo")

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
