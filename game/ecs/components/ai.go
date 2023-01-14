package components

import (
	"encoding/json"

	game "github.com/co0p/tankism/game/ecs"
	"github.com/co0p/tankism/lib/ecs"
)

const AIType = "AI"

type AI struct {
	TargetGroup int
}

func (AI) Type() ecs.ComponentType {
	return AIType
}

func (AI) Load(data []byte) game.LoadableComponent {
	var nt AI
	json.Unmarshal(data, &nt)
	return nt
}
