package game

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init() error
	Draw(image *ebiten.Image)
	HandleInput()
	Update() error
}

type GameScene struct {
	Systems       []ecs.System
	EntityManager ecs.EntityManager
}

func (s *GameScene) Init() error {
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	for _, v := range s.Systems {
		v.Draw(screen)
	}
}

func (s *GameScene) Update() error {
	var err error
	for _, v := range s.Systems {
		err = v.Update()
	}
	return err
}

func (s *GameScene) HandleInput() {}
