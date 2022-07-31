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

	entities := s.EntityManager.FindByComponents(components.SpriteType, components.TranslateType)

	// we sort the slice by z-index
	sort.Slice(entities, func(a, b int) bool {
		spriteA := entities[a].GetComponent(components.SpriteType).(*components.Sprite)
		spriteB := entities[b].GetComponent(components.SpriteType).(*components.Sprite)
		return spriteA.ZIndex < spriteB.ZIndex
	})

	// now draw them
	for _, e := range entities {

		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		img := sprite.Image
		rect := img.Bounds()

		x := translate.X
		y := translate.Y
		rotation := translate.Rotation
		//scale := translate.Scale

		op := &ebiten.DrawImageOptions{}
		//op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(float64(rect.Dx())/2, float64(rect.Dy())/2)
		op.GeoM.Translate(x, y)

		screen.DrawImage(img, op)
	}
}
