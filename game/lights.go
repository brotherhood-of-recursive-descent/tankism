package game

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

var DEFAULT_COLOR = color.RGBA{255, 255, 255, 255} // white

func NewAmbientLight(e *ecs.Entity) {
	c := &components.AmbientLight{
		Color: color.RGBA{128, 128, 128, 255},
	}
	e.AddComponent(c)
}

func NewPointLight(e *ecs.Entity, x, y float64) {
	i, err := media.LoadImage(media.LightPoint)
	img := ebiten.NewImageFromImage(i)
	if err != nil {
		panic("expected LightPoint media to exist")
	}

	light := &components.Light{Image: img, Color: DEFAULT_COLOR}
	translate := &components.Translate{
		X:        x,
		Y:        y,
		Scale:    2,
		Rotation: 0,
	}

	e.AddComponents(light, translate)
}

func NewCircleLight(e *ecs.Entity, x, y float64) {
	i, err := media.LoadImage(media.LightCircle)
	img := ebiten.NewImageFromImage(i)
	if err != nil {
		panic("expected LightPoint media to exist")
	}

	light := &components.Light{Image: img, Color: DEFAULT_COLOR}
	translate := &components.Translate{
		X:        x,
		Y:        y,
		Scale:    3,
		Rotation: 0,
	}

	e.AddComponents(light, translate)
}