package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

const TransformType = "transform"

// Transform holds all information needed to position the entity in the world
type Transform struct {
	Point vector.Vec2d

	Scale    float64
	Rotation float64

	Parent  *Transform
	OffsetX float64
	OffsetY float64

	Children []*Transform
}

func (t *Transform) AddParent(p *Transform) {
	t.Parent = p
	p.addChild(t)
}

func (t *Transform) addChild(c *Transform) {
	t.Children = append(t.Children, c)
}

func (t Transform) Type() ecs.ComponentType {
	return TransformType
}
