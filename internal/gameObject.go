package internal

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameObject interface {
	OnUpdate(*World)
	OnDraw(*sdl.Renderer)
	OnInput(sdl.KeyboardEvent)
}

type Position struct {
	X, Y int
	R    float32
}

type Movement struct {
	Vx, Vy float32
}
