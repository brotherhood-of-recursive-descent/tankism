package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

const CameraType = "Camera"

type Camera struct {
	Point vector.Vec2d
}

func (t Camera) Type() ecs.ComponentType {
	return CameraType
}
