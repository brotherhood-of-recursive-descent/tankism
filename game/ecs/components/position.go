package components

import "github.com/co0p/tankism/lib/ecs"

const TranslateType = "translate"

// Translate holds all information needed to position the entity in the world
type Translate struct {
	X float64
	Y float64

	Scale    float64
	Rotation float64
}

func (t Translate) Type() ecs.ComponentType {
	return TranslateType
}
