package singleplayer

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

func configureTank(tank *ecs.Entity) {

	img, _ := media.LoadImage(media.TankImage)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s}
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
		Rotation: 0.05,
	}

	tank.AddComponents(sprite, translate, velocity, shaking, controller)
}

func configureFpsCounter(fps *ecs.Entity, width int) {
	fps.AddComponent(&components.Text{
		Value: "0",
		Font:  media.FontMedium,
		Color: lib.ColorGreen})
	fps.AddComponent(&components.Translate{
		X:        float64(width - 120),
		Y:        50.0,
		Scale:    1,
		Rotation: 0.0,
	})
	fps.AddComponent(&components.FPS{})
}
