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

func configureLight(e *ecs.Entity, mode ebiten.CompositeMode, X, Y float64) {

	sprite, _ := media.LoadImage(media.LightCircle)

	e.AddComponent(&components.Light{
		Image:         ebiten.NewImageFromImage(sprite),
		CompositeMode: mode,
		Scale:         1,
	})
	e.AddComponent(&components.Translate{
		X: X,
		Y: Y,
	})
}

func configureBigTank(e *ecs.Entity) {

	t, _ := media.LoadImage(media.BigTankImage)
	s := ebiten.NewImageFromImage(t)

	l, _ := media.LoadImage(media.LightCircle)
	lightSprite := ebiten.NewImageFromImage(l)

	sprite := &components.Sprite{Image: s}
	shaking := &components.Shaking{}
	translate := &components.Translate{
		X:     700.0,
		Y:     700.0,
		Scale: 1,
	}

	light := &components.Light{
		Image:         ebiten.NewImageFromImage(lightSprite),
		CompositeMode: ebiten.CompositeModeDestinationIn,
		Scale:         1,
	}

	e.AddComponents(sprite, translate, shaking, light)
}
