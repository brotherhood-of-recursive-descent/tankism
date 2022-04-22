package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const FPSType = "FPS"

// FPS, empty component filtering
type FPS struct {
}

func (t FPS) Type() ecs.ComponentType {
	return FPSType
}
