package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const MapType = "Map"

type Map struct {
	// use later for effects/modifiers applied to moveable objects
}

func (t Map) Type() ecs.ComponentType {
	return MapType
}
