package game

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
)

type GameSceneWithCamera struct {
	Systems       []ecs.System
	EntityManager ecs.EntityManager
	Camera        camera.Camera // TODO move to own camera wrapping this

	game Game
}

func (s *GameSceneWithCamera) Init() error {
	return nil
}

func (s *GameSceneWithCamera) Draw(screen *ebiten.Image) {

	s.Camera.Surface.Clear()

	for _, v := range s.Systems {
		v.Draw(s.Camera.Surface)
	}

	s.Camera.Blit(screen)

}

func (s *GameSceneWithCamera) Update() error {
	var err error
	for _, v := range s.Systems {
		err = v.Update()
	}
	return err
}

func (s *GameSceneWithCamera) HandleInput() {}
