package components_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
)

var goldenTargetComponent = components.Target{GroupId: 23}
var goldenTargetJson = `{"GroupId":23}`

func Test_Target_Serialize(t *testing.T) {
	// when
	bytez, err := json.Marshal(goldenTargetComponent)

	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if string(bytez) != goldenTargetJson {
		t.Errorf("expected export to match: \ngot: %s\nwant: %s", string(bytez), goldenTargetJson)
	}
}

func Test_Target_Deserialize(t *testing.T) {

	// given some bytes
	input := []byte(goldenTargetJson)

	// when
	var actual components.Target
	err := json.Unmarshal(input, &actual)

	// then we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if !reflect.DeepEqual(goldenTargetComponent, actual) {
		t.Errorf("expected deserialzed to equal acutual.\ngot: %v\nwant:%v\n", actual, goldenTargetComponent)
	}
}
