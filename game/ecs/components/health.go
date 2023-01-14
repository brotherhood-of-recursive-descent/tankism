package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const HealthType = "Health"

type Health struct {
	Value int
}

func (t Health) Type() ecs.ComponentType {
	return HealthType
}
