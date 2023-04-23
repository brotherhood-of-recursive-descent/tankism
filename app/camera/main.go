package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/camera"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type CameraDemo struct {
	game.GameSceneWithCamera // type embedding

	game            *game.Game
	cameraComponent *components.Camera // to control the camera
}

func (s *CameraDemo) Init() error {

	w, h := s.game.WindowSize()
	s.Camera = *camera.NewCamera(w, h)
	s.cameraComponent = &components.Camera{
		Zoom:       1.0,
		CameraMode: camera.CameraModeDefault,
	}

	s.Systems = append(s.Systems,
		&systems.SpriteRenderer{EntityManager: &s.EntityManager},
		&systems.Controller{EntityManager: &s.EntityManager},
		&systems.MovementSystem{EntityManager: &s.EntityManager},
		&systems.TextRenderer{EntityManager: &s.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &s.EntityManager},

		systems.NewCameraSystem(&s.EntityManager, &s.Camera),
	)

	// the objects int the world
	fps := s.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	tank := s.EntityManager.NewEntity()
	game.NewTankWithPosition(tank, 400, 400)
	tank.AddComponent(s.cameraComponent)

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

func (s *CameraDemo) HandleInput() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.cameraComponent.Zoom += 0.01
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.cameraComponent.Zoom -= 0.01
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.cameraComponent.Zoom = 1
	}
}

func main() {

	game := game.NewGame()
	demo := CameraDemo{game: game}

	game.AddScene("Demo", &demo)
	game.SetScene("Demo")

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(800, 600)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
