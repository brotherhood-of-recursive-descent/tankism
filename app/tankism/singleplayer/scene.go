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
		&systems.LightingSystem{
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

	l1 := s.entityManager.NewEntity()
	configureLight(l1, ebiten.CompositeModeSourceOver, 100, 500)

	l2 := s.entityManager.NewEntity()
	configureLight(l2, ebiten.CompositeModeClear, 300, 500)

	l3 := s.entityManager.NewEntity()
	configureLight(l3, ebiten.CompositeModeDestination, 500, 500)

	l4 := s.entityManager.NewEntity()
	configureLight(l4, ebiten.CompositeModeDestinationOver, 700, 500)

	l5 := s.entityManager.NewEntity()
	configureLight(l5, ebiten.CompositeModeSourceIn, 900, 500)

	l6 := s.entityManager.NewEntity()
	configureLight(l6, ebiten.CompositeModeDestinationIn, 1100, 500)

	bigTank := s.entityManager.NewEntity()
	configureBigTank(bigTank)

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
