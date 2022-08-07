package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const CleanupType = "Cleanup"

// Cleanup, empty component filtering
type Cleanup struct {
}

func (t Cleanup) Type() ecs.ComponentType {
	return CleanupType
}
