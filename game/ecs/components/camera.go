package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const CameraType = "Camera"

type CameraMode int

const (
	CameraModeDefault = iota
	CameraModeCenter
)

type Camera struct {
	CameraMode CameraMode
	Zoom       float64
}

func (t Camera) Type() ecs.ComponentType {
	return CameraType
}
