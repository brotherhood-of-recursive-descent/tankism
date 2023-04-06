package components

import (
	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

const RotationType = "Rotation"

type Rotation struct {
	Point    vector.Vec2d
	Rotation float64
}

func (t Rotation) Type() ecs.ComponentType {
	return RotationType

}
