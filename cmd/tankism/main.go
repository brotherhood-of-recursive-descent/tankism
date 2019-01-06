package main

import (
	"log"
	"time"

	"github.com/co0p/tankism/internal/objects"

	"github.com/co0p/tankism/internal"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {

	ttf.Init()

	windowManager := internal.NewSDLWindowManager(600, 600, "Tankism")
	window, err := windowManager.CreateWindow()
	defer window.Destroy()
	if err != nil {
		log.Fatal("creating window: ", err)
	}

	assetManager, err := internal.NewAssetManager("assets")
	if err != nil {
		log.Fatal("asset manager: ", err)
	}

	font, err := ttf.OpenFont("assets/fonts/test.ttf", 24)
	if err != nil {
		log.Fatal("open font: ", err)
	}

	renderer, err := windowManager.GetRenderer()
	world := internal.World{}

	tankSprite := assetManager.GetTexture(renderer, "tank_dark.png")
	tank := objects.NewPlayerTank(tankSprite)
	fpsCounter := objects.NewFPSCounter(&world, renderer, font)

	world.Add(fpsCounter)
	world.Add(tank)

	running := true
	for running {

		// event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				world.UpdateInput(*e)
			case *sdl.QuitEvent:
				log.Println("Quit")
				running = false
				break
			}
		}

		world.UpdatePhysics()
		world.UpdateDraw(renderer)

		time.Sleep(10 * time.Millisecond)
	}
	sdl.Quit()
}
