package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type TextRenderer struct {
	EntityManager *ecs.EntityManager
}

func (s *TextRenderer) Update() error {
	return nil
}

func (s *TextRenderer) Draw(screen *ebiten.Image) {

	entities := s.EntityManager.FindByComponents(components.TextType, components.TransformType)

	for _, e := range entities {

		op := &ebiten.DrawImageOptions{}

		translate := e.GetComponent(components.TransformType).(*components.Transform)
		x := translate.X
		y := translate.Y
		scale := translate.Scale

		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(x, y)

		txt := e.GetComponent(components.TextType).(*components.Text)

		text.Draw(screen, txt.Value, txt.Font, int(x), int(y), txt.Color)
	}

}
