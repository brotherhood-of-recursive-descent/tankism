package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Controller struct {
	EntityManager *ecs.EntityManager
}

func (s *Controller) Draw(screen *ebiten.Image) {}

func (s *Controller) Update() error {

	shooters := s.EntityManager.FindByComponents(components.TransformType, components.ShootingType)
	for _, e := range shooters {
		shooting := e.GetComponent(components.ShootingType).(*components.Shooting)

		if shooting.Cooldown < 0 && ebiten.IsKeyPressed(ebiten.KeySpace) {
			shooting.Active = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		entities := s.EntityManager.FindByComponents(components.AmbientLightType)
		if len(entities) == 1 {
			ambientLight := entities[0].GetComponent(components.AmbientLightType).(*components.AmbientLight)
			ambientLight.Active = !ambientLight.Active
		}
	}

	return nil
}
