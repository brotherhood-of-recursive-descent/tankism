package main

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
}

func NewFPSCounter() *FPSCounter {

	font, err := ttf.OpenFont("assets/fonts/test.ttf", 24)
	if err != nil {
		log.Fatal("open font: ", err)
	}

	return &FPSCounter{
		currentCount: 0,
		intervalTime: time.Now(),
		font:         font,
		fpsString:    " ",
	}
}

func (fpsc *FPSCounter) Update() {

	fpsc.currentCount++
	t := time.Now()
	if t.After(fpsc.intervalTime.Add(1 * time.Second)) {
		fpsc.intervalTime = time.Now()
		fpsc.fpsString = fmt.Sprintf("FPS: %d", fpsc.currentCount)
		fpsc.currentCount = 0
	}
}

func (fpsc *FPSCounter) Draw(renderer *sdl.Renderer) {

	// TODO: only create new texture when necessary

	FPSsurface, err := fpsc.font.RenderUTF8Solid(fpsc.fpsString, sdl.Color{0, 255, 0, 200})
	if err != nil {
		log.Fatal("render font surface: ", err)
	}

	FPStexture, err := renderer.CreateTextureFromSurface(FPSsurface)
	if err != nil {
		log.Fatal("create texture: ", err)
	}

	renderer.Copy(FPStexture, &sdl.Rect{X: 0, Y: 0, W: 120, H: 26}, &sdl.Rect{X: 10, Y: 10, W: 120, H: 26})

	defer FPSsurface.Free()
	defer FPStexture.Destroy()
}

func main() {

	ttf.Init()

	windowManager := internal.NewSDLWindowManager(600, 600, "Tankism")
	window, err := windowManager.CreateWindow()
	defer window.Destroy()

	assetManager, err := internal.NewAssetManager("assets")
	if err != nil {
		log.Fatal("asset manager: ", err)
	}

	renderer, err := windowManager.GetRenderer()
	tankTexture := assetManager.GetTexture(renderer, "tank_dark.png")

	fpsCounter := NewFPSCounter()

	running := true
	for running {

		// event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN && e.Keysym.Sym == sdl.K_f {
					log.Println("toggle Fullscreen")
					windowManager.ToogleFullscreen()
				}
			case *sdl.QuitEvent:
				log.Println("Quit")
				running = false
				break
			}
		}

		// logic
		fpsCounter.Update()

		// drawing
		renderer.SetDrawColor(255, 244, 233, 255)
		renderer.Clear()
		renderer.Copy(tankTexture, &sdl.Rect{X: 0, Y: 0, W: 42, H: 46}, &sdl.Rect{X: windowManager.GetWidth()/2 - 21, Y: windowManager.GetHeight()/2 - 23, W: 42, H: 46})
		fpsCounter.Draw(renderer)
		renderer.Present()

		// make sure we don't use all cpu power
		//time.Sleep(20 * time.Millisecond)
	}
	sdl.Quit()
}
