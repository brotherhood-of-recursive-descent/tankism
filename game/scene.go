package game

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Systems       []ecs.System
	EntityManager ecs.EntityManager
}

func (s *Scene) Init(sm *lib.SceneManager) error {
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, v := range s.Systems {
		v.Draw(screen)
	}
}

func (s *Scene) Update() error {
	var err error
	for _, v := range s.Systems {
		err = v.Update()
	}
	return err
}

func (s *Scene) HandleInput() {}
