package game

import (
	"encoding/json"
	"time"

	"github.com/co0p/tankism/lib/ecs"
)

type State struct {
	Name       string
	ModifiedAt time.Time
	Entities   []ecs.Entity
}

func NewState(data []byte) (State, error) {

	var state State
	err := json.Unmarshal(data, &state)

	return state, err
}

func (s *State) Save(name string) ([]byte, error) {
	now := time.Now()
	s.ModifiedAt = now
	s.Name = name
	return json.Marshal(s)
}
