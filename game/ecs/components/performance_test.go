package components_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/co0p/tankism/game/ecs/components"
)

var goldenPerformanceComponent = components.Performance{
	ShowFPS:         true,
	ShowTPS:         true,
	ShowEntityCount: true,
}

var goldenPerformanceJson = `{"ShowFPS":true,"ShowTPS":true,"ShowEntityCount":true,"ShowGraph":false,"HistoryLength":0,"PastFPS":null}`

func Test_Performance_Serialize(t *testing.T) {

	// when
	bytez, err := json.Marshal(goldenPerformanceComponent)

	// then  we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	// and export should match golden string
	if string(bytez) != goldenPerformanceJson {
		t.Errorf("expected export to match: \ngot: %s\nwant: %s", string(bytez), goldenPerformanceJson)
	}
}

func Test_Performance_Deserialize(t *testing.T) {

	// given some bytes
	input := []byte(goldenPerformanceJson)

	// when
	var actual components.Performance
	err := json.Unmarshal(input, &actual)

	// then we have no error
	if err != nil {
		t.Errorf("expected err to be nil, got %v\n", err)
	}

	// and export should match golden string
	if !reflect.DeepEqual(goldenPerformanceComponent, actual) {
		t.Errorf("expected deserialzed to equal acutual.\ngot: %v\nwant:%v\n", actual, goldenPerformanceComponent)
	}
}
