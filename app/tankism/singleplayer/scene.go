package singleplayer

import (
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SinglePlayerScene struct {
	WindowWidth  int
	WindowHeight int

	sceneManager *lib.SceneManager

	entityManager *ecs.EntityManager
	systems       []ecs.System
}

func NewSinglePlayerScene(sceneManager *lib.SceneManager) *SinglePlayerScene {

	entityManager := ecs.EntityManager{}

	// empty scene
	scene := SinglePlayerScene{}

	var s []ecs.System
	s = append(s,
		&systems.SpriteRenderer{
			EntityManager: &entityManager,
		},
		&systems.Shaker{
			EntityManager: &entityManager,
		},
		&systems.TextRenderer{
			EntityManager: &entityManager,
		},
		&systems.PerformanceMonitor{
			EntityManager: &entityManager,
		},
		&systems.Controller{
			EntityManager: &entityManager,
		},
	)

	scene.entityManager = &entityManager
	scene.systems = s

	return &scene
}

func (s *SinglePlayerScene) Init(sm *lib.SceneManager) error {

	tank := s.entityManager.NewEntity()
	configureTank(tank)

	fpsCounter := s.entityManager.NewEntity()
	configureFpsCounter(fpsCounter, sm.ScreenWidth)

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
