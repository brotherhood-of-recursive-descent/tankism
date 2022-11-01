package ecs_test

import (
	"encoding/json"
	"testing"

	"github.com/co0p/tankism/lib/ecs"
)

var goldenEntityJson = `{"Test1":{"Var1":42,"Var2":42.32},"Test2":{"Var1":"some","Var2":["where","over"]}}`

type TestComponent struct {
	Var1 int
	Var2 float64
}

func (TestComponent) Type() ecs.ComponentType { return ecs.ComponentType("Test1") }

type TestComponent2 struct {
	Var1 string
	Var2 []string
}

func (TestComponent2) Type() ecs.ComponentType { return ecs.ComponentType("Test2") }

func Test_Entity_Serialize_to_json(t *testing.T) {

	// given
	em := ecs.NewEntityManager(nil)
	e := em.NewEntity()
	e.AddComponents(&TestComponent{Var1: 42, Var2: 42.32}, &TestComponent2{Var1: "some", Var2: []string{"where", "over"}})

	// when
	bytez, err := json.Marshal(e)

	// then
	if err != nil {
		t.Errorf("expect err to be nil, got %s\n", err)
	}

	actualString := string(bytez)
	if goldenEntityJson != actualString {
		t.Errorf("expected strings to match, \ngot: %s\nwant: %s\n", actualString, goldenEntityJson)
	}
}

func Test_Entity_Deserialize_from_json(t *testing.T) {
	// todo
}
