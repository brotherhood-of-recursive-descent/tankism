package main

import (
	_ "embed"
	"log"

	"github.com/co0p/tankism/app/tankism/exit"
	"github.com/co0p/tankism/app/tankism/menu"
	"github.com/co0p/tankism/app/tankism/singleplayer"
	"github.com/co0p/tankism/app/tankism/start"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Client struct {
	windowWidth  int
	windowHeight int

	sceneManager *lib.SceneManager

	/*
		audioMixer audio.Mixer
		loader media.Loader
	*/
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
	sceneManager := lib.NewSceneManager()

	menuScene := menu.NewMenuScene(sceneManager)
	startScene := start.NewStartScreen(sceneManager)
	exitScene := exit.NewExitScene(sceneManager)
	singlePlayerScene := singleplayer.NewSinglePlayerScene(sceneManager)

	sceneManager.RegisterScene("MENU", menuScene)
	sceneManager.RegisterScene("START", startScene)
	sceneManager.RegisterScene("EXIT", exitScene)
	sceneManager.RegisterScene("SINGLEPLAYER", singlePlayerScene)

	sceneManager.ChangeScene("START")

	client := &Client{}
	client.sceneManager = sceneManager

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
