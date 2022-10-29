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

	entities := s.EntityManager.FindByComponents(components.SpriteType, components.TransformType)

	// we sort the slice by z-index
	sort.Slice(entities, func(a, b int) bool {
		spriteA := entities[a].GetComponent(components.SpriteType).(*components.Sprite)
		spriteB := entities[b].GetComponent(components.SpriteType).(*components.Sprite)
		return spriteA.ZIndex < spriteB.ZIndex
	})

	// now draw them
	for _, e := range entities {
		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		img := sprite.Image
		if img == nil {
			continue
		}
		rect := img.Bounds()

		translate := e.GetComponent(components.TransformType).(*components.Transform)
		x := translate.X
		y := translate.Y
		rotation := translate.Rotation

		op := &ebiten.DrawImageOptions{}
		//op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(float64(rect.Dx())/2, float64(rect.Dy())/2)
		op.GeoM.Translate(x, y)

		// scale
		// rotate
		// translate

		screen.DrawImage(img, op)
	}
}
