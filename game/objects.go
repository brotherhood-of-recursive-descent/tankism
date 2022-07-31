package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewDrum(e *ecs.Entity, x, y float64) {

	img, _ := media.LoadImage(media.BarrelGray)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Translate{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate)
}

func NewCrate(e *ecs.Entity, x, y float64) {

	img, _ := media.LoadImage(media.CrateWood)
	s := ebiten.NewImageFromImage(img)

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Translate{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate)
}
