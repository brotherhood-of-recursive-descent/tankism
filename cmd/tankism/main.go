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

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

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
		renderer.Present()

		// make sure we don't use all cpu power
		time.Sleep(20 * time.Millisecond)
	}
	sdl.Quit()
}
