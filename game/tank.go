package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewTank(tank *ecs.Entity) {

	img, _ := media.LoadImage(media.TankImage)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	velocity := &components.Velocity{
		IntertiaMax: 2,
		Rotation:    0.05,
	}
	shaking := &components.Shaking{}
	controller := &components.Controller{}
	translate := &components.Translate{
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
