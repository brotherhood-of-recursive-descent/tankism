package internal

import "github.com/veandco/go-sdl2/sdl"

type World struct {
	gameObjects []GameObject
}

func (w *World) UpdatePhysics() {
	for _, o := range w.gameObjects {
		o.OnUpdate(w)
	}
}

func (w *World) UpdateInput(e sdl.KeyboardEvent) {
	for _, o := range w.gameObjects {
		o.OnInput(e)
	}
}

func (w *World) UpdateDraw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	for _, o := range w.gameObjects {
		o.OnDraw(renderer)
	}

	renderer.Present()
}

func (w *World) Add(o GameObject) {
	w.gameObjects = append(w.gameObjects, o)
}
