package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const TriggerType = "Trigger"

// Trigger, empty component filtering
type Trigger struct {
	Action func(e *ecs.Entity, em *ecs.EntityManager)
}

func (t Trigger) Type() ecs.ComponentType {
	return TriggerType
}
