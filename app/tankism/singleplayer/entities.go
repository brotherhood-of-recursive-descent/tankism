package singleplayer

import (
	"image/color"
	"math/rand"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/media"
	"github.com/hajimehoshi/ebiten/v2"
)

func configureTank(tank *ecs.Entity) {

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
		Color:         color.RGBA{R: 10, G: 10, B: 0, A: 160},
	}
	e.AddComponents(ambientLight)
}

type Tilemap [][]string

func configureMap(e *ecs.Entity, tilemap Tilemap, w int, h int) {
	mapImage := ebiten.NewImage(w, h)
	t1, _ := media.LoadImage(media.TileGrassImage1)
	tileImage1 := ebiten.NewImageFromImage(t1)
	t2, _ := media.LoadImage(media.TileGrassImage2)
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
	translate := &components.Translate{
		Scale: 1,
	}

	e.AddComponents(sprite, translate)
}
