package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const DamageType = "Damage"

type Damage struct {
	Value int
}

func (t Damage) Type() ecs.ComponentType {
	return DamageType
}
