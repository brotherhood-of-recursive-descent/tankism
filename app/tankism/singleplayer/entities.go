package singleplayer

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

func configureTank(tank *ecs.Entity) {

	img, _ := media.LoadImage(media.TankImage)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 1}
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

func configureFpsCounter(fps *ecs.Entity, width int) {
	fps.AddComponent(&components.Text{
		Value: "0",
		Font:  media.FontMedium,
		Color: lib.ColorGreen,
	})
	fps.AddComponent(&components.Translate{
		X:        float64(width - 120),
		Y:        50.0,
		Scale:    1,
		Rotation: 0.0,
	})
	fps.AddComponent(&components.FPS{})
}

func configureAITank(e *ecs.Entity) {

	t, _ := media.LoadImage(media.BigTankImage)
	s := ebiten.NewImageFromImage(t)

	sprite := &components.Sprite{Image: s}
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
		Color:         color.RGBA{R: 10, G: 10, B: 0, A: 160},
	}
	e.AddComponents(ambientLight)
}
