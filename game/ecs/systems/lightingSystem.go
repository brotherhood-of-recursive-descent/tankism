package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type LightingSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *LightingSystem) Update() error { return nil }
func (s *LightingSystem) Draw(screen *ebiten.Image) {

	entities := s.EntityManager.FindByComponents(components.LightType, components.TranslateType)

	for _, e := range entities {

		light := e.GetComponent(components.LightType).(*components.Light)
		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		rect := light.Image.Bounds()

		x := translate.X
		y := translate.Y
		rotation := light.Rotation
		scale := light.Scale

		op := &ebiten.DrawImageOptions{}
		op.CompositeMode = light.CompositeMode
		op.GeoM.Translate(-float64(rect.Dx())/2, -float64(rect.Dy())/2)
		op.GeoM.Scale(scale, scale)
		op.GeoM.Rotate(rotation)
		op.GeoM.Translate(x, y)

		img := light.Image
		screen.DrawImage(img, op)
	}

}
