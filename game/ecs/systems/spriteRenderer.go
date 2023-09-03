package systems

import (
	"sort"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderer struct {
	EntityManager *ecs.EntityManager
}

func (s *SpriteRenderer) Update() error { return nil }
func (s *SpriteRenderer) Draw(screen *ebiten.Image) {

	// for rendering we need a sprite and a target position
	entities := s.EntityManager.FindByComponents(components.SpriteType, components.TransformType)

	// we sort them by z-index to maintain the visual ordering
	sort.Slice(entities, func(a, b int) bool {
		spriteA := entities[a].GetComponent(components.SpriteType).(*components.Sprite)
		spriteB := entities[b].GetComponent(components.SpriteType).(*components.Sprite)
		return spriteA.ZIndex < spriteB.ZIndex
	})

	// now we draw them
	for _, e := range entities {
		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		img := sprite.Image
		if img == nil {
			continue
		}

		// TODO: maybe add w,h to sprite component to avoid fetching bounds
		rect := img.Bounds()
		width := float64(rect.Dx())
		height := float64(rect.Dy())
		translate := e.GetComponent(components.TransformType).(*components.Transform)
		scaleFactor := translate.Scale

		// we understand rotation along the center of the image,
		// therefore we move the image, rotate, move back and then reposition it.
		// This deviates from the normal way which is scale, rotate and then translate
		op := &ebiten.DrawImageOptions{}

		// first scale
		op.GeoM.Scale(scaleFactor, scaleFactor)

		// second rotate
		op.GeoM.Translate(-width*scaleFactor/2, -height*scaleFactor/2)
		op.GeoM.Rotate(translate.Rotation)
		op.GeoM.Translate(width*scaleFactor/2, height*scaleFactor/2)

		// third transalte
		op.GeoM.Translate(translate.Point.X, translate.Point.Y)

		screen.DrawImage(img, op)
	}
}
