package components

import "github.com/co0p/tankism/lib/ecs"

const TranslateType = "transform"

// Transform holds all information needed to position the entity in the world
type Transform struct {
	X float64
	Y float64

	Scale    float64
	Rotation float64

	Parent   *Transform
	Children []*Transform
	OffsetX  float64
	OffsetY  float64
}

func (t *Transform) AddParent(p *Transform) {
	t.Parent = p
	p.addChild(t)
}

func (t *Transform) addChild(c *Transform) {
	t.Children = append(t.Children, c)
}

func (t Transform) Type() ecs.ComponentType {
	return TranslateType
}
