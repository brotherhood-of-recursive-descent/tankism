package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const CollisionType = "Collision"

// Collision, omponent for filtering collision
type Collision struct {
	Target *ecs.Entity
}

func (t Collision) Type() ecs.ComponentType {
	return CollisionType
}
