package game

import (
	"image/color"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/resource"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	DEFAULT_COLOR = color.RGBA{255, 255, 255, 255} // white
	MUTED_COLOR   = color.RGBA{128, 128, 128, 128} // gray
)

func NewAmbientLight(e *ecs.Entity) {
	c := &components.AmbientLight{
		Color: MUTED_COLOR,
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
	NewCircleLightWithColor(e, x, y, color.White)
}

func NewCircleLightWithColor(e *ecs.Entity, x, y float64, clr color.Color) {
	i, err := media.LoadImage(media.LightCircle)
	img := ebiten.NewImageFromImage(i)
	if err != nil {
		panic("expected LightPoint media to exist")
	}

	light := &components.Light{Image: img, Color: clr}
	translate := &components.Translate{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(light, translate)
}

func NewLightSpritesheet() resource.SpriteSheet {
	spritesheet, err := resource.NewSpriteSheet(media.Lights, 128, 128)
	if err != nil {
		panic("failed to load light spritesheet: " + err.Error())
	}

	return spritesheet
}
