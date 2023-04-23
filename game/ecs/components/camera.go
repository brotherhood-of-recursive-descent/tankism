package components

import (
	"github.com/co0p/tankism/lib/camera"
	"github.com/co0p/tankism/lib/ecs"
)

const CameraType = "Camera"

// Camera is a component to control the camera
type Camera struct {
	CameraMode camera.CameraMode
	Zoom       float64
}

func (t Camera) Type() ecs.ComponentType {
	return CameraType
}
