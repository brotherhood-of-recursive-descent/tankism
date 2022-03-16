package main

import (
	_ "embed"
	"github.com/co0p/tankism/app/tankism/menu"
	"github.com/co0p/tankism/app/tankism/start"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Client struct {
	windowWidth  int
	windowHeight int
	startScene   lib.Scene
	menuScene    lib.Scene

	sceneManager *lib.SceneManager

	/* scenes scenes.Manager
	audioMixer audio.Mixer
	loader media.Loader */
}

func (t *Client) Draw(screen *ebiten.Image) {
	t.sceneManager.Draw(screen)
}

func (t *Client) Update() error {
	return t.sceneManager.Update()
}

func (t *Client) Layout(outsideWidth, outsideHeight int) (int, int) {
	t.windowWidth = outsideWidth
	t.windowHeight = outsideHeight

	t.sceneManager.SetWindowDimension(t.windowWidth, t.windowHeight)

	return outsideWidth, outsideHeight
}

func NewClient() *Client {
	sceneManager := &lib.SceneManager{}

	menuScene := menu.NewMenuScene(sceneManager)
	startScene := start.NewStartScreen(sceneManager, menuScene)

	sceneManager.ChangeScene(startScene)

	client := &Client{}
	client.sceneManager = sceneManager
	client.startScene = startScene
	client.menuScene = menuScene

	return client
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
