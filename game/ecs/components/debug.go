package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const DebugType = "Debug"

type Debug struct{}

func (Debug) Type() ecs.ComponentType {
	return DebugType
}
