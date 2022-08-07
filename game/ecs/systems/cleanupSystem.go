package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CleanupSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *CleanupSystem) Draw(screen *ebiten.Image) {}

func (s *CleanupSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.CleanupType)

	for _, e := range entities {
		s.EntityManager.RemoveEntity(e)
	}
	return nil
}
