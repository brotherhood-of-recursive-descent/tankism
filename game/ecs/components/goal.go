package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const GoalType = "Goal"

type Goal struct{}

func (t Goal) Type() ecs.ComponentType {
	return GoalType
}
