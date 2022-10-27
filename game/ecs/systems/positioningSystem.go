package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type PositioningSysystem struct {
	EntityManager *ecs.EntityManager
}

func (s *PositioningSysystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.TranslateType)

	for _, e := range entities {
		translate := e.GetComponent(components.TranslateType).(*components.Transform)

		// no parent means that it is potential root node
		if translate.Parent == nil {
			// TODO(#36) - implement walking the graph
			for _, child := range translate.Children {
				child.X = translate.X + child.OffsetX
				child.Y = translate.Y + child.OffsetY
			}
		}
	}

	return nil
}

func (s *PositioningSysystem) Draw(screen *ebiten.Image) {}
