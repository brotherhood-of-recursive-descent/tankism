package ecs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	libecs "github.com/co0p/tankism/game/ecs"
	"github.com/co0p/tankism/lib/ecs"
)

var Test1Type ecs.ComponentType = "Test1"

type TestComponent struct {
	Var1 int
	Var2 float64
}

func (TestComponent) Load(data []byte) libecs.LoadableComponent {
	var nt TestComponent
	json.Unmarshal(data, &nt)
	return nt
}

func (TestComponent) Type() ecs.ComponentType { return Test1Type }

var Test2Type ecs.ComponentType = "Test2"

type TestComponent2 struct {
	Var1 string
	Var2 []string
}

func (TestComponent2) Load(data []byte) libecs.LoadableComponent {
	var nt TestComponent2
	json.Unmarshal(data, &nt)
	return nt
}

func (TestComponent2) Type() ecs.ComponentType { return Test2Type }

func Test_EntityLoader_Load(t *testing.T) {

	// given an entity
	em := ecs.EntityManager{}
	e := em.NewEntity()
	e.AddComponents(
		&TestComponent{Var1: 42, Var2: 42.42},
		&TestComponent2{Var1: "fourtytwo", Var2: []string{"fourty", "two"}},
	)
	json, _ := json.Marshal(e)

	// when loading
	loader := libecs.NewEntityLoader()
	loader.Register(TestComponent{}, TestComponent2{})
	fmt.Printf("got '%s' to load... \n", json)

	target := em.NewEntity()
	err := loader.Load(target, json)

	// we have the no error and the data
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if !target.HasComponent(Test1Type) {
		t.Errorf("expected entity to have test1type\n")
	}

	if !target.HasComponent(Test2Type) {
		t.Errorf("expected entity to have test2type\n")
	}

	testcmp1 := target.GetComponent(Test1Type).(TestComponent)
	if testcmp1.Var1 != 42 {
		t.Errorf("expected TestComponent to Var1 = 42\n")
	}
}
