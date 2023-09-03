package systems

import (
	"math"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// MotionControlSystem controls the motion of the main player.
type MotionControlSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *MotionControlSystem) Draw(screen *ebiten.Image) {}

func (s *MotionControlSystem) Update() error {

	entities := s.EntityManager.FindByComponents(components.MotionControlType, components.TransformType)

	for _, e := range entities {

		control := e.GetComponent(components.MotionControlType).(*components.MotionControl)
		transform := e.GetComponent(components.TransformType).(*components.Transform)
		newRotation := transform.Rotation

		// we want the object to move along a direction vector:
		// https://gamedev.stackexchange.com/questions/130830/how-to-calcolate-moving-vector-of-object-by-its-rotation-degress-in-2d
		if ebiten.IsKeyPressed(control.KeyForward) {
			transform.Point.X -= control.AccelerationRate * (math.Sin(newRotation-math.Pi) * 2)
			transform.Point.Y += control.AccelerationRate * (math.Cos(newRotation-math.Pi) * 2)
			control.Accelerate()
		}

		if ebiten.IsKeyPressed(control.KeyBackward) {
			transform.Point.X += control.AccelerationRate * (math.Sin(newRotation-math.Pi) * 2)
			transform.Point.Y -= control.AccelerationRate * (math.Cos(newRotation-math.Pi) * 2)
			control.Accelerate()
		}

		if ebiten.IsKeyPressed(control.KeyRight) {
			// adjust the rotation directly on the transform, avoiding the element to spin
			transform.Rotation += control.RotationRate
		}

		if ebiten.IsKeyPressed(control.KeyLeft) {
			// adjust the rotation directly on the transform, avoiding the element to spin
			transform.Rotation -= control.RotationRate
		}

		if inpututil.IsKeyJustReleased(control.KeyForward) || inpututil.IsKeyJustReleased(control.KeyBackward) {
			control.AccelerationRate = 0
		}
	}

	return nil
}
