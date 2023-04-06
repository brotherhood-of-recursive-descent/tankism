package systems

import (
	"math"
	"math/rand"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type ShootingSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *ShootingSystem) Draw(screen *ebiten.Image) {}

func (s *ShootingSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.ShootingType, components.TransformType)

	for _, e := range entities {
		shooting := e.GetComponent(components.ShootingType).(*components.Shooting)
		transform := e.GetComponent(components.TransformType).(*components.Transform)

		// system should fire!
		if shooting.Active {
			bullet := s.EntityManager.NewEntity()
			game.NewBullet(bullet, 0, 0)
			bullet.AddComponent(&components.Transform{
				Point: vector.Vec2d{
					X: transform.Point.X + 10,
					Y: transform.Point.Y + 10,
				},
				Rotation: transform.Rotation + 2*math.Pi,
			})
			bullet.AddComponent(&components.Velocity{
				X: 1.1,
				Y: 1.1,
			})
			bullet.AddComponent(&components.Damage{
				Value: 100,
			})

			shooting.Active = false
			shooting.Cooldown = shooting.CooldownMin + rand.Intn(shooting.CooldownMax-shooting.CooldownMin+1)
		} else {
			shooting.Cooldown--
		}

	}
	return nil
}
