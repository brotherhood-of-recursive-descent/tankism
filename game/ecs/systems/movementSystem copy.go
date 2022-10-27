package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *MovementSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.VelocityType, components.TranslateType)

	for _, e := range entities {
		velocity := e.GetComponent(components.VelocityType).(*components.Velocity)
		translate := e.GetComponent(components.TranslateType).(*components.Transform)
		translate.X = translate.X + velocity.X
		translate.Y = translate.Y + velocity.Y
		translate.Rotation = translate.Rotation + velocity.Rotation
	}

	return nil
}

func (s *MovementSystem) Draw(screen *ebiten.Image) {}
