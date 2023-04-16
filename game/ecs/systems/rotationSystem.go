package systems

import (
	"fmt"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type RotationSystem struct {
	EntityManager *ecs.EntityManager
}

func (s *RotationSystem) Draw(screen *ebiten.Image) {}

func (s *RotationSystem) Update() error {

	entities := s.EntityManager.FindByComponents(components.RotationType, components.TransformType)

	for _, e := range entities {
		rotation := e.GetComponent(components.RotationType).(*components.Rotation)
		transform := e.GetComponent(components.TransformType).(*components.Transform)

		// first move rotation-center to origin
		origin := vector.Vec2d{}
		moved := origin.Subtract(rotation.Point)

		// apply the rotation
		rotated := moved.Rotate(transform.Rotation)

		// move it back
		movedBack := rotated.Add(rotation.Point)

		// and finally store it in the transform
		newTransform := transform.Point.Add(movedBack)
		fmt.Printf("transform: %v\nrotation: %v\nnew: %v\n\n", transform.Point, rotation.Point, newTransform)
		transform.Point = newTransform

	}

	return nil
}
