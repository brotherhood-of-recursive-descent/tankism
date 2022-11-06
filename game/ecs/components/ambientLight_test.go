package components_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

var goldenAmbientLightComponent = components.AmbientLight{CompositeMode: ebiten.CompositeModeCopy, Color: lib.ColorRed, Active: true}
var goldenAmbientLightJson = `{"CompositeMode":2,"Color":{"R":250,"G":37,"B":37,"A":255},"Active":true}`

func Test_AmbientLight_Serialize(t *testing.T) {
	// when
	bytez, err := json.Marshal(goldenAmbientLightComponent)

	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if string(bytez) != goldenAmbientLightJson {
		t.Errorf("expected export to match: \ngot: %s\nwant: %s", string(bytez), goldenAmbientLightJson)
	}
}

func Test_AmbientLight_Deserialize(t *testing.T) {

	// given some bytes
	input := []byte(goldenAmbientLightJson)

	// when
	var actual components.AmbientLight
	err := json.Unmarshal(input, &actual)

	// then we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	if !reflect.DeepEqual(goldenAmbientLightComponent, actual) {
		t.Errorf("expected deserialzed to equal acutual.\ngot: %v\nwant:%v\n", actual, goldenAmbientLightComponent)
	}
}
