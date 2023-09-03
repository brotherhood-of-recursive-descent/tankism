package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type MotionSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *MotionSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.MotionType, components.TransformType)

	for _, e := range entities {
		motion := e.GetComponent(components.MotionType).(*components.Motion)
		translate := e.GetComponent(components.TransformType).(*components.Transform)

		translate.Point = translate.Point.Add(motion.Velocity)
		translate.Rotation += motion.AngularVelocity
	}

	return nil
}

func (s *MotionSystem) Draw(screen *ebiten.Image) {}
