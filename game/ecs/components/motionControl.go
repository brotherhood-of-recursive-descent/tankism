package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const defaultAccelerationRate float64 = 0.05
const defaultRotationRate float64 = 0.05
const maxAcceleration float64 = 2

const MotionControlType = "MotionControl"

// MotionControl contains the keymapping and other settings related to motion
type MotionControl struct {
	KeyForward  ebiten.Key
	KeyBackward ebiten.Key
	KeyRight    ebiten.Key
	KeyLeft     ebiten.Key

	AccelerationRate float64
	RotationRate     float64
}

func (t MotionControl) Type() ecs.ComponentType {
	return MotionControlType
}

func (t *MotionControl) Accelerate() {
	if t.AccelerationRate == 0 {
		t.AccelerationRate = defaultAccelerationRate
	}
	if t.AccelerationRate < maxAcceleration {
		t.AccelerationRate += defaultAccelerationRate
	}
}

func NewMotionControl() *MotionControl {
	return &MotionControl{
		KeyForward:  ebiten.KeyArrowDown,
		KeyBackward: ebiten.KeyArrowUp,
		KeyLeft:     ebiten.KeyArrowLeft,
		KeyRight:    ebiten.KeyArrowRight,

		AccelerationRate: defaultAccelerationRate,
		RotationRate:     defaultRotationRate,
	}
}
