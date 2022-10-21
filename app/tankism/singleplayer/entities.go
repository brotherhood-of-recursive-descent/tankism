package singleplayer

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

func configureTank(tank *ecs.Entity) {

	img, _ := resources.LoadImage(resources.TankImage)
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

func configureAITank(e *ecs.Entity) {

	t, _ := resources.LoadImage(resources.BigTankImage)
	s := ebiten.NewImageFromImage(t)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Translate{
		X:     700.0,
		Y:     700.0,
		Scale: 1,
	}
	ai := &components.AI{
		TargetGroup: 1,
	}

	e.AddComponents(sprite, translate, ai)
}

func configureAmbientLight(e *ecs.Entity) {

	ambientLight := &components.AmbientLight{
		CompositeMode: ebiten.CompositeModeSourceOver,
		Color:         color.RGBA{R: 60, G: 76, B: 128, A: 10},
	}
	e.AddComponents(ambientLight)
}
