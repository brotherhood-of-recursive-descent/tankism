package singleplayer

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

type SinglePlayerScene struct {
	WindowWidth  int
	WindowHeight int

	sceneManager *lib.SceneManager

	entityManager ecs.EntityManager
	systems       []ecs.System
}

func NewSinglePlayerScene(sceneManager *lib.SceneManager) *SinglePlayerScene {

	// load media
	img, _ := media.LoadImage(media.TankImage)
	sprite := ebiten.NewImageFromImage(img)

	// build the tank
	entityManager := ecs.EntityManager{}
	tank := entityManager.NewEntity()

	tank.AddComponent(&components.Sprite{Image: sprite})
	tank.AddComponent(&components.Translate{
		X:        200.0,
		Y:        200.0,
		Scale:    1,
		Rotation: 0.10,
	})

	var s []ecs.System
	s = append(s,
		&systems.SpriteRenderer{
			EntityManager: entityManager,
		},
		&systems.Shaker{
			EntityManager: entityManager,
		},
	)

	scene := SinglePlayerScene{
		entityManager: entityManager,
		systems:       s,
	}

	return &scene
}

func (s *SinglePlayerScene) Init() error {
	// load tank image

	return nil
}

func (s *SinglePlayerScene) Draw(screen *ebiten.Image) {
	for _, v := range s.systems {
		v.Draw(screen)
	}
}

func (s *SinglePlayerScene) Update() error {
	var err error
	for _, v := range s.systems {
		err = v.Update()
	}
	return err
}

func (s *SinglePlayerScene) WindowDimension() (int, int) {
	return s.WindowWidth, s.WindowHeight
}

func (s *SinglePlayerScene) SetWindowDimension(w, h int) {
	s.WindowWidth = w
	s.WindowHeight = h
}
