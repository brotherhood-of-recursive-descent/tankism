package objects

import (
	"log"

	"github.com/co0p/tankism/internal"
	"github.com/veandco/go-sdl2/sdl"
)

type Tank struct {
	X, Y, R int32
	sprite  *sdl.Texture
}

func NewPlayerTank(sprite *sdl.Texture) *Tank {
	return &Tank{
		X:      200/2 - 21,
		Y:      200/2 - 21,
		R:      0,
		sprite: sprite,
	}
}

func (t *Tank) OnUpdate(w *internal.World) {}

func (t *Tank) OnDraw(r *sdl.Renderer) {
	src := &sdl.Rect{X: 0, Y: 0, W: 42, H: 46}
	dest := &sdl.Rect{X: t.X, Y: t.Y, W: 42, H: 46}
	r.Copy(t.sprite, src, dest)
}

func (t *Tank) OnInput(e sdl.KeyboardEvent) {

	if e.Type == sdl.KEYDOWN {
		switch e.Keysym.Sym {
		case sdl.K_w:
			log.Println("up")
			t.Y = t.Y + 5
		case sdl.K_s:
			log.Println("down")
			t.Y = t.Y - 5
		case sdl.K_a:
			log.Println("left")
			t.X = t.X - 5
		case sdl.K_d:
			log.Println("right")
			t.X = t.X + 5
		}
	}
}
