package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const ShakingType = "Shaking"

// Shaking, empty component for filtering shakables
type Shaking struct {
}

func (t Shaking) Type() ecs.ComponentType {
	return ShakingType
}
