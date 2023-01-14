package components

import (
	"encoding/json"

	game "github.com/co0p/tankism/game/ecs"
	"github.com/co0p/tankism/lib/ecs"
)

const AimingType = "AIAiming"

type Aiming struct {
	TargetGroup int
}

func (Aiming) Type() ecs.ComponentType {
	return AimingType
}

func (Aiming) Load(data []byte) game.LoadableComponent {
	var nt Aiming
	json.Unmarshal(data, &nt)
	return nt
}
