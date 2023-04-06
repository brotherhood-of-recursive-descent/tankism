package systems_test

import (
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

func Test_RotationSystem_identity(t *testing.T) {

	// given

	givenTransform := components.Transform{
		Point: vector.Vec2d{
			X: 100,
			Y: 100,
		},
		Rotation: 100,
	}
	nullRotation := components.Rotation{
		Point: vector.Vec2d{
			X: 0, Y: 0,
		},
		Rotation: 0,
	}
	em := ecs.NewEntityManager(nil)
	s := systems.RotationSystem{EntityManager: em}

	e1 := em.NewEntity()
	e1.AddComponents(
		givenTransform,
		nullRotation,
	)

	// when
	s.Update()

	// then
	actualTransform := e1.GetComponent(components.TransformType).(components.Transform)

	if actualTransform.Point.X != givenTransform.Point.X ||
		actualTransform.Point.Y != givenTransform.Point.Y ||
		actualTransform.Rotation != givenTransform.Rotation {
		t.Errorf("expected transform to be %v, got %v\n", givenTransform, actualTransform)
	}

}

// get all entities with
// - translateComponent
// - rotationComponent

// calculate new x,y based on rotationComponent

// save in new rotation and X,Y in translate
