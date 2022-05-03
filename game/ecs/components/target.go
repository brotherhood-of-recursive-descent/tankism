package components

import "github.com/co0p/tankism/lib/ecs"

const TargetType = "Target"

type Target struct {
	GroupId int
}

func (Target) Type() ecs.ComponentType {
	return TargetType
}
