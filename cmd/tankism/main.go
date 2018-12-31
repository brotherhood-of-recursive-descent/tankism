package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("Tankism", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		log.Println("creating window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		// event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
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
