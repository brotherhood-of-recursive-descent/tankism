package systems

import (
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

	// todo, sort by z-index to maintain right order of rendering
	for _, e := range entities {

		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		rect := sprite.Image.Bounds()

		x := translate.X
		y := translate.Y
		rotation := translate.Rotation
		scale := translate.Scale

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(x, y)

		img := sprite.Image
		screen.DrawImage(img, op)
	}

}
