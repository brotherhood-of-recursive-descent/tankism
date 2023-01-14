package components_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
)

var goldenAIComponent = components.Aiming{TargetGroup: 42}
var goldenAIJson = `{"TargetGroup":42}`

func Test_AI_Serialize(t *testing.T) {
	// when
	bytez, err := json.Marshal(goldenAIComponent)

	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if string(bytez) != goldenAIJson {
		t.Errorf("expected export to match: \ngot: %s\nwant: %s", string(bytez), goldenAIJson)
	}
}

func Test_AI_Deserialize(t *testing.T) {

	// given some bytes
	input := []byte(goldenAIJson)

	// when
	var actual components.Aiming
	err := json.Unmarshal(input, &actual)

	// then we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if !reflect.DeepEqual(goldenAIComponent, actual) {
		t.Errorf("expected deserialzed to equal acutual.\ngot: %v\nwant:%v\n", actual, goldenAIComponent)
	}
}
