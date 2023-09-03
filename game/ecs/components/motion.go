package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

const MotionType = "Motion"

// Motion stores the current velocity as well as angular velocity
type Motion struct {
	Velocity        vector.Vec2d
	AngularVelocity float64
}

func (s Motion) Type() ecs.ComponentType {
	return MotionType
}
