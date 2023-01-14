package game

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewDrum(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.BarrelGray)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}

	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate, bbox)
}

func NewCrate(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.CrateWood)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate, bbox)
}

func NewCrateMetal(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.CrateMetal)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}
	goal := &components.Goal{}

	e.AddComponents(sprite, translate, bbox, goal)
}

func NewTree(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.BigTreeImage)
	s := ebiten.NewImageFromImage(img)
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate)
}

func NewBullet(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.BulletSandOutline)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate, bbox)
}

func NewOilSpill(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.OilSpillLarge)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate, bbox)
}

func NewBarricade(e *ecs.Entity, x, y float64) {

	img, _ := resources.LoadImage(resources.Barricade)
	s := ebiten.NewImageFromImage(img)
	w, h := s.Size()
	bbox := &components.BoundingBox{
		Width:  float64(w),
		Height: float64(h),
	}
	sprite := &components.Sprite{Image: s, ZIndex: 100}
	translate := &components.Transform{
		X:        x,
		Y:        y,
		Scale:    1,
		Rotation: 0,
	}

	e.AddComponents(sprite, translate, bbox)
}
