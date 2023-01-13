package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionResolution struct {
	EntityManager *ecs.EntityManager
}

func (s *CollisionResolution) Draw(screen *ebiten.Image) {}

func (s *CollisionResolution) Update() error {

	entities := s.EntityManager.FindByComponents(components.CollisionType)

	for _, e := range entities {
		collision := e.GetComponent(components.CollisionType).(*components.Collision)
		// bounce back -- basic resolution for all
		// TODO resolve based on other types e.g. Health
		if collision.Target.HasComponent(components.VelocityType) {
			v := collision.Target.GetComponent(components.VelocityType).(*components.Velocity)
			v.Intertia = -1
		}
		e.RemoveComponent(components.CollisionType)
	}

	return nil
}
