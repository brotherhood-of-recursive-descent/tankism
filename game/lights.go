package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	DEFAULT_COLOR = lib.Color{255, 255, 255, 255} // white
	MUTED_COLOR   = lib.Color{128, 128, 128, 128} // gray
)

func NewAmbientLight(e *ecs.Entity) {
	c := &components.AmbientLight{
		Color: MUTED_COLOR,
	}
	e.AddComponent(c)
}

func NewPointLight(e *ecs.Entity, x, y float64) {
	i, err := resources.LoadImage(resources.LightPoint)
	img := ebiten.NewImageFromImage(i)
	if err != nil {
		panic("expected LightPoint media to exist")
	}

	light := &components.Light{Image: img, Color: DEFAULT_COLOR}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    2,
		Rotation: 0,
	}

	e.AddComponents(light, translate)
}

func NewCircleLight(e *ecs.Entity, x, y float64) {
	NewCircleLightWithColor(e, x, y, lib.ColorWhite)
}

func NewCircleLightWithColor(e *ecs.Entity, x, y float64, clr lib.Color) {
	i, err := resources.LoadImage(resources.LightCircle)
	img := ebiten.NewImageFromImage(i)
	if err != nil {
		panic("expected LightPoint media to exist")
	}

	light := &components.Light{Image: img, Color: clr}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    2,
		Rotation: 0,
	}

	e.AddComponents(light, translate)
}

func NewLightSpritesheet() resource.SpriteSheet {
	spritesheet, err := resource.NewSpriteSheet(resources.Lights, 128, 128)
	if err != nil {
		panic("failed to load light spritesheet: " + err.Error())
	}

	return spritesheet
}
