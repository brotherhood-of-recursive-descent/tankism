package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const ControllerType = "Controller"

type Controller struct{}

func (t Controller) Type() ecs.ComponentType {
	return ControllerType
}
