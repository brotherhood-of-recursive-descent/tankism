package objects

import (
	"fmt"
	"log"
	"time"

	"github.com/co0p/tankism/internal"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type FPSCounter struct {
	currentCount int
	intervalTime time.Time
	font         *ttf.Font
	fpsString    string
	isVisible    bool
}

func NewFPSCounter(world *internal.World, renderer *sdl.Renderer, font *ttf.Font) *FPSCounter {
	return &FPSCounter{
		currentCount: 0,
		intervalTime: time.Now(),
		font:         font,
		fpsString:    " ",
		isVisible:    false,
	}
}

func (o *FPSCounter) OnInput(e sdl.KeyboardEvent) {
	if e.Type == sdl.KEYDOWN && e.Keysym.Sym == sdl.K_p {
		o.isVisible = !o.isVisible
	}
}
func (o *FPSCounter) OnUpdate(w *internal.World) {

	o.currentCount++
	t := time.Now()
	if t.After(o.intervalTime.Add(1 * time.Second)) {
		o.intervalTime = time.Now()
		o.fpsString = fmt.Sprintf("FPS: %d", o.currentCount)
		o.currentCount = 0
	}
}

func (o *FPSCounter) OnDraw(r *sdl.Renderer) {
	if !o.isVisible {
		return
	}

	FPSsurface, err := o.font.RenderUTF8Solid(o.fpsString, sdl.Color{R: 0, G: 255, B: 0, A: 200})
	if err != nil {
		log.Fatal("render font surface: ", err)
	}

	FPStexture, err := r.CreateTextureFromSurface(FPSsurface)
	if err != nil {
		log.Fatal("create texture: ", err)
	}

	r.Copy(FPStexture, &sdl.Rect{X: 0, Y: 0, W: 120, H: 26}, &sdl.Rect{X: 10, Y: 10, W: 120, H: 26})

	defer FPSsurface.Free()
	defer FPStexture.Destroy()
}
