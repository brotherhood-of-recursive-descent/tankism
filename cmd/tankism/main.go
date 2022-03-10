package main

import (
	_ "embed"
	"github.com/co0p/tankism"
	"github.com/co0p/tankism/cmd/tankism/start"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Client struct {
	windowWidth  int
	windowHeight int
	startScene   tankism.Scene

	/* scenes scenes.Manager
	audioMixer audio.Mixer
	loader assets.Loader */
}

func (t *Client) Draw(screen *ebiten.Image) {
	t.startScene.Draw(screen)
}

func (t *Client) Update() error {
	return t.startScene.Update()
}

func (t *Client) Layout(outsideWidth, outsideHeight int) (int, int) {
	t.windowWidth = outsideWidth
	t.windowHeight = outsideHeight

	t.startScene.SetWindowDimension(t.windowWidth, t.windowHeight)

	return outsideWidth, outsideHeight
}

func NewClient() *Client {

	return &Client{
		startScene: start.NewStartScreen(),
	}
}

func main() {

	// setup
	ebiten.SetFullscreen(true)

	// construct game stuff
	client := NewClient()

	// start the game
	if err := ebiten.RunGame(client); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
