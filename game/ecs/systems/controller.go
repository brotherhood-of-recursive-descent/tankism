package systems

import (
	"math"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var Alpha uint8 = 0

type Controller struct {
	EntityManager *ecs.EntityManager
}

func (s *Controller) Draw(screen *ebiten.Image) {}

func (s *Controller) Update() error {

	entities := s.EntityManager.FindByComponents(components.ControllerType, components.TransformType, components.VelocityType)

	for _, e := range entities {

		translate := e.GetComponent(components.TransformType).(*components.Transform)
		velocity := e.GetComponent(components.VelocityType).(*components.Velocity)
		newRotation := translate.Rotation

		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			translate.X -= velocity.Intertia * (math.Sin(newRotation-math.Pi) * 2)
			translate.Y += velocity.Intertia * (math.Cos(newRotation-math.Pi) * 2)
			velocity.IncreaseInertia()
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			translate.X += velocity.Intertia * (math.Sin(newRotation-math.Pi) * 2)
			translate.Y -= velocity.Intertia * (math.Cos(newRotation-math.Pi) * 2)
			velocity.IncreaseInertia()
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			translate.Rotation += velocity.Rotation
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			translate.Rotation -= velocity.Rotation
		}

		if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) || inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
			velocity.ResetInertia()
		}

	}

	shooters := s.EntityManager.FindByComponents(components.ControllerType, components.TransformType, components.ShootingType)
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
