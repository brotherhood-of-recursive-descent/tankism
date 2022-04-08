package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderer struct {
	EntityManager ecs.EntityManager
}

func (s *SpriteRenderer) Update() error { return nil }
func (s *SpriteRenderer) Draw(screen *ebiten.Image) {

	entities := s.EntityManager.FindByComponents(components.SpriteType, components.TranslateType)

	// todo, sort by z-index to maintain right order of rendering
	for _, e := range entities {

		op := &ebiten.DrawImageOptions{}

		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		x := translate.X
		y := translate.Y
		scale := translate.Scale

		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(x, y)

		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		img := sprite.Image
		screen.DrawImage(img, op)
	}

}
