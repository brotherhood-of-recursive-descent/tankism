package systems_test

import (
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
)

func TestPositioningSystem(t *testing.T) {

	em := ecs.EntityManager{}
	system := systems.PositioningSysystem{EntityManager: &em}

	parent := em.NewEntity()
	parentTransform := components.Transform{X: 111.11, Y: 222.22}
	parent.AddComponent(&parentTransform)

	child := em.NewEntity()
	childTransform := components.Transform{OffsetX: 222.22, OffsetY: 333.33}
	childTransform.AddParent(&parentTransform)
	child.AddComponent(&childTransform)

	err := system.Update()
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if childTransform.X != 333.33 || childTransform.Y != 555.55 {
		t.Errorf("expected child transform to be (x:%v, y:%v), got %v\n", 333.33, 555.55, childTransform)

	}
}
