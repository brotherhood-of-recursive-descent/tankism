package systems

import (
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type PositioningSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *PositioningSystem) Update() error {
	entities := s.EntityManager.FindByComponents(components.TransformType)

	for _, e := range entities {
		node := e.GetComponent(components.TransformType).(*components.Transform)

		// always start at the root
		if isRoot(node) {
			applyPosition(node, node.Children)
		}
	}

	return nil
}

func isRoot(n *components.Transform) bool {
	return n.Parent == nil
}

func applyPosition(p *components.Transform, cs []*components.Transform) {
	for _, child := range cs {
		child.X = p.X + child.OffsetX
		child.Y = p.Y + child.OffsetY

		if len(child.Children) > 0 {
			applyPosition(child, child.Children)
		}
	}

}

func (s *PositioningSystem) Draw(screen *ebiten.Image) {}
