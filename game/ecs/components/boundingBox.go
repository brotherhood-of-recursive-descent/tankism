package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const BoundingBoxType = "BoundingBox"

// BoundingBox, component for collision detection
type BoundingBox struct {
	Width  int
	Height int
}

func (t BoundingBox) Type() ecs.ComponentType {
	return BoundingBoxType
}
