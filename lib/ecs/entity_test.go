package ecs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/co0p/tankism/lib/ecs"
)

var goldenEntityJson = `{"Test1":{"Var1":42,"Var2":42.32},"Test2":{"Var1":"some","Var2":["where","over"]}}`

var Test1Type ecs.ComponentType = "Test1"

type TestComponent struct {
	Var1 int
	Var2 float64
}

func (TestComponent) Type() ecs.ComponentType { return Test1Type }

var Test2Type ecs.ComponentType = "Test2"

type TestComponent2 struct {
	Var1 string
	Var2 []string
}

func (TestComponent2) Type() ecs.ComponentType { return Test2Type }

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

func XTest_Entity_Deserialize_from_json(t *testing.T) {

	em := ecs.NewEntityManager(nil)
	e := em.NewEntity()

	data := []byte(goldenEntityJson)
	var dto map[ecs.ComponentType]interface{}
	err := json.Unmarshal(data, &dto)

	if err != nil {
		t.Errorf("expect err to be nil, got %s\n", err)
	}

	for k, m := range dto {
		switch k {
		case Test1Type:
			e.AddComponent(m.(TestComponent))
		case Test2Type:
			e.AddComponent(m.(TestComponent2))
		default:
			fmt.Println("unknown type")
		}
	}

	test1 := e.GetComponent(Test1Type).(*TestComponent)
	if test1.Var1 != 42 {
		t.Errorf("expected test1 component to have value 42, got %d\n", test1.Var1)
	}
}
