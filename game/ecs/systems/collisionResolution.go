package systems

import (
	"fmt"

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

		// collision with goal
		if collision.Target.HasComponent(components.GoalType) {
			fmt.Println("you won!")
			collision.Target.AddComponent(&components.Cleanup{})
		}

		// damage giver collided with destructable entity
		if e.HasComponent(components.DamageType) && collision.Target.HasComponent(components.HealthType) {
			damage := e.GetComponent(components.DamageType).(*components.Damage)
			health := collision.Target.GetComponent(components.HealthType).(*components.Health)

			health.Value = health.Value - damage.Value

			if health.Value <= 0 {
				collision.Target.AddComponent(&components.Cleanup{})
			}

			e.AddComponent(&components.Cleanup{})
		}

		// bounce back -- basic resolution for all
		// TODO resolve based on other types e.g. Health
		if collision.Target.HasComponent(components.VelocityType) {
			v := collision.Target.GetComponent(components.VelocityType).(*components.Velocity)
			v.Intertia = -1
		}

		// collission resolved
		e.RemoveComponent(components.CollisionType)
	}

	return nil
}
