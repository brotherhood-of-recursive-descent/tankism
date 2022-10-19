package game

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/sound"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Systems       []ecs.System
	EntityManager ecs.EntityManager
	SoundManager  *sound.SoundManager
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) Init(sm *lib.SceneManager) error {
	s.SoundManager = &sm.SoundManager
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
