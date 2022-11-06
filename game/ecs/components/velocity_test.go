package components_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
)

var goldenVelocityComponent = components.Velocity{
	X:           1,
	Y:           2,
	Rotation:    100,
	IntertiaMax: 200,
	Intertia:    300,
}

var goldenVelocityJson = `{"X":1,"Y":2,"Rotation":100,"IntertiaMax":200,"Intertia":300}`

func Test_Velocity_Serialize(t *testing.T) {

	// when
	bytez, err := json.Marshal(goldenVelocityComponent)

	// then  we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	// and export should match golden string
	if string(bytez) != goldenVelocityJson {
		t.Errorf("expected export to match: \ngot: %s\nwant: %s", string(bytez), goldenVelocityJson)
	}
}

func Test_Velocity_Deserialize(t *testing.T) {

	// given some bytes
	input := []byte(goldenVelocityJson)

	// when
	var actual components.Velocity
	err := json.Unmarshal(input, &actual)

	// then we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	// and export should match golden string
	if !reflect.DeepEqual(goldenVelocityComponent, actual) {
		t.Errorf("expected deserialzed to equal acutual.\ngot: %v\nwant:%v\n", actual, goldenVelocityComponent)
	}
}
