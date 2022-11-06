package components

import (
	"github.com/co0p/tankism/lib/ecs"
)

const PerformanceType = "Performance"

// Performance
type Performance struct {
	ShowFPS         bool
	ShowTPS         bool
	ShowEntityCount bool

	ShowGraph     bool
	HistoryLength int
	PastFPS       []float64
}

func (t Performance) Type() ecs.ComponentType {
	return PerformanceType
}
