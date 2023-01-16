package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const CameraType = "Camera"

type Camera struct {
	X float64
	Y float64
}

func (t Camera) Type() ecs.ComponentType {
	return CameraType
}
