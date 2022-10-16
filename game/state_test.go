package game_test

import (
	"fmt"
	"testing"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/test"
)

func Test_StateLoad_should_load_game_state(t *testing.T) {

	state, err := game.NewState(test.GameState_Valid)

	if err != nil {
		t.Errorf("expected err to be nil, got %s\n", err)
	}

	if len(state.Entities) != 1 {
		t.Errorf("expected one entity, got %d\n", len(state.Entities))
		fmt.Printf("%v\n", state)

	}
}

func Test_StateSave_should_save_game_state(t *testing.T) {

	state, _ := game.NewState(test.GameState_Valid)

	res, err := state.Save("some_name")
	if err != nil {
		t.Errorf("expected err to be nil, got %d\n", err)
	}

	if len(res) == 0 {
		t.Errorf("expected res not to be empty\n")
	}

	loadedState, _ := game.NewState(res)
	if loadedState.Name != "some_name" {
		t.Errorf("expected Name to be 'some_name', got %s\n", loadedState.Name)
	}
}
