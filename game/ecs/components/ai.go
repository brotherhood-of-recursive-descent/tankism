package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const AIType = "AI"

type AI struct{}

func (t AI) Type() ecs.ComponentType {
	return AIType
}
