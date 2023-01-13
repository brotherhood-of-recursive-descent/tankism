package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewTank(tank *ecs.Entity) {

	img, _ := resources.LoadImage(resources.TankImage)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	velocity := &components.Velocity{
		IntertiaMax: 2,
		Rotation:    0.05,
	}
	shaking := &components.Shaking{}
	controller := &components.Controller{}
	translate := &components.Transform{
		X:        200.0,
		Y:        200.0,
		Scale:    1,
		Rotation: 0,
	}
	target := &components.Target{
		GroupId: 1,
	}

	tank.AddComponents(controller, sprite, translate, velocity, shaking, target)
}

func NewBigTank(tank *ecs.Entity, x float64, y float64) {

	img, _ := resources.LoadImage(resources.BigTankImage)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	shaking := &components.Shaking{}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	tank.AddComponents(sprite, translate, shaking)
}
