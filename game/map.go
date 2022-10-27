package game

import (
	"math/rand"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap [][]string

func NewMap(e *ecs.Entity, tilemap Tilemap, w int, h int) {
	mapImage := ebiten.NewImage(w, h)
	t1, _ := resources.LoadImage(resources.TileGrassImage1)
	tileImage1 := ebiten.NewImageFromImage(t1)
	t2, _ := resources.LoadImage(resources.TileGrassImage2)
	tileImage2 := ebiten.NewImageFromImage(t2)
	tileW, tileH := tileImage1.Size()

	for y := 0; y < h; y += tileH {
		for x := 0; x < w; x += tileW {
			tileImage := tileImage1
			if rand.Intn(2) > 0 {
				tileImage = tileImage2
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			mapImage.DrawImage(tileImage, op)
		}
	}

	sprite := &components.Sprite{Image: mapImage, ZIndex: 0}
	translate := &components.Transform{
		Scale: 1,
	}

	e.AddComponents(sprite, translate)
}
