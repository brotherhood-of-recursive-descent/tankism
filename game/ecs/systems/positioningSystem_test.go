package systems_test

import (
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

func TestPositioningSystem(t *testing.T) {

	em := ecs.EntityManager{}
	system := systems.PositioningSystem{EntityManager: &em}

	parent := em.NewEntity()
	parentTransform := components.Transform{Point: vector.Vec2d{X: 111.11, Y: 222.22}}
	parent.AddComponent(&parentTransform)

	child := em.NewEntity()
	childTransform := components.Transform{OffsetX: 222.22, OffsetY: 333.33}
	childTransform.AddParent(&parentTransform)
	child.AddComponent(&childTransform)

	err := system.Update()
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if childTransform.Point.X != 333.33 || childTransform.Point.Y != 555.55 {
		t.Errorf("expected child transform to be (x:%v, y:%v), got %v\n", 333.33, 555.55, childTransform)

	}
}

func TestPositioningSystem_multiple_level(t *testing.T) {

	/*
	     A
	     | \
	     B  C
	    / \
	   D   E
	*/
	em := ecs.EntityManager{}
	system := systems.PositioningSystem{EntityManager: &em}

	a := em.NewEntity()
	aTransform := components.Transform{Point: vector.Vec2d{X: 1, Y: 0}}
	a.AddComponent(&aTransform)

	b := em.NewEntity()
	bTransform := components.Transform{OffsetX: 2}
	bTransform.AddParent(&aTransform)
	b.AddComponent(&bTransform)

	c := em.NewEntity()
	cTransform := components.Transform{OffsetX: 3}
	cTransform.AddParent(&aTransform)
	c.AddComponent(&cTransform)

	d := em.NewEntity()
	dTransform := components.Transform{OffsetX: 4}
	dTransform.AddParent(&bTransform)
	d.AddComponent(&dTransform)

	e := em.NewEntity()
	eTransform := components.Transform{OffsetX: 5}
	eTransform.AddParent(&bTransform)
	e.AddComponent(&eTransform)

	// WHEN relative position calculation
	err := system.Update()
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	// THEN
	if aTransform.Point.X != 1 {
		t.Errorf("expected a.x to be 1, got %v\n", aTransform.Point.X)
	}

	if bTransform.Point.X != 3 {
		t.Errorf("expected b.x to be 3, got %v\n", bTransform.Point.X)
	}

	if cTransform.Point.X != 4 {
		t.Errorf("expected c.x to be 4, got %v\n", cTransform.Point.X)
	}

	if dTransform.Point.X != 7 {
		t.Errorf("expected a.x to be 7, got %v\n", dTransform.Point.X)
	}

	if eTransform.Point.X != 8 {
		t.Errorf("expected a.x to be 8, got %v\n", aTransform.Point.X)
	}
}
