package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const ShootingType = "Shooting"

type Shooting struct {
	CooldownMin int
	CooldownMax int
	Cooldown    int

	Active bool
}

func (t Shooting) Type() ecs.ComponentType {
	return ShootingType
}
