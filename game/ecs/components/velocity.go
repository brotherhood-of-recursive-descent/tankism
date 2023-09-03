package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const VelocityType = "velocity"

const increaseFactor = 0.05

// Deprecated: please use the Motion component instead.
type Velocity struct {
	X float64
	Y float64

	Rotation float64

	IntertiaMax float64
	Intertia    float64
}

func (v *Velocity) IncreaseInertia() {
	if v.Intertia == 0 {
		v.Intertia = increaseFactor
	}
	if v.Intertia < v.IntertiaMax {
		v.Intertia += increaseFactor
	}
}

func (v *Velocity) ResetInertia() {
	v.Intertia = 0
}

func (Velocity) Type() ecs.ComponentType {
	return VelocityType
}
