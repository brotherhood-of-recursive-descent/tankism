package state_test

import (
	"fmt"
	"testing"

	"github.com/co0p/tankism/game/state"
	"github.com/co0p/tankism/test"
)

func Test_New_should_load_game_state(t *testing.T) {

	s, err := state.New(test.GameState_Valid)

	if err != nil {
		t.Errorf("expected err to be nil, got %s\n", err)
	}

	if len(s.Entities) != 1 {
		t.Errorf("expected one entity, got %d\n", len(s.Entities))
		fmt.Printf("%v\n", s)

	}
}

func Test_Save_should_save_game_state(t *testing.T) {

	s, _ := state.New(test.GameState_Valid)

	res, err := s.Save("some_name")
	if err != nil {
		t.Errorf("expected err to be nil, got %d\n", err)
	}

	if len(res) == 0 {
		t.Errorf("expected res not to be empty\n")
	}

	loadedState, _ := state.New(res)
	if loadedState.Name != "some_name" {
		t.Errorf("expected Name to be 'some_name', got %s\n", loadedState.Name)
	}
}
