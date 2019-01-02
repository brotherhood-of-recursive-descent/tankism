package main

import (
	"log"
	"time"

	"github.com/co0p/tankism/internal"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	windowManager := internal.NewSDLWindowManager(600, 600, "Tankism")
	window, err := windowManager.CreateWindow()
	defer window.Destroy()

	assetManager, err := internal.NewAssetManager("assets")
	if err != nil {
		log.Fatal("asset manager: ", err)
	}

	renderer, err := windowManager.GetRenderer()
	tankTexture := assetManager.GetTexture(renderer, "tank_dark.png")

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

		// drawing
		renderer.SetDrawColor(255, 244, 233, 255)
		renderer.Clear()
		renderer.Copy(tankTexture, &sdl.Rect{X: 0, Y: 0, W: 42, H: 46}, &sdl.Rect{X: windowManager.GetWidth()/2 - 21, Y: windowManager.GetHeight()/2 - 23, W: 42, H: 46})
		renderer.Present()

		// make sure we don't use all cpu power
		time.Sleep(20 * time.Millisecond)
	}
	sdl.Quit()
}
