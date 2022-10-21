package game

import (
	"fmt"
	"os"

	"github.com/co0p/tankism/lib/sound"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is the shell of a game and is a meant to be used as a container for scenes
type Game struct {
	Resources  resources.ResourceManager
	SoundMixer sound.Mixer

	// scene management
	scenes       map[string]Scene
	currentScene Scene
	nextScene    Scene

	// game stuff
	initialized  bool
	ScreenWidth  int
	ScreenHeight int
}

func NewGame() *Game {
	return &Game{
		scenes: make(map[string]Scene),
	}
}

func (g *Game) Exit() {
	fmt.Println("Game::Exit")
	os.Exit(0)
}

func (g *Game) WindowSize() (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

// ebiten stuff
func (g *Game) Draw(screen *ebiten.Image) {

	if g.currentScene == nil {
		return
	}
	g.currentScene.Draw(screen)
}

func (g *Game) Update() error {

	// lazy initialization
	if !g.initialized {
		for _, i := range g.scenes {
			if err := i.Init(); err != nil {
				return err
			}
		}
		fmt.Println("Game::Initialized done")
		g.initialized = true
	}

	// switch to next scene if set
	if g.currentScene != g.nextScene && g.nextScene != nil {
		g.currentScene = g.nextScene
		g.nextScene = nil
	}

	if g.currentScene == nil {
		return nil
	}

	err := g.currentScene.Update()
	g.currentScene.HandleInput()
	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.ScreenWidth = outsideWidth
	g.ScreenHeight = outsideHeight
	return g.ScreenWidth, g.ScreenHeight
}

// scene stuff
func (g *Game) AddScene(sceneKey string, scene Scene) {
	g.scenes[sceneKey] = scene
}

func (g *Game) SetScene(handle string) {
	scene, ok := g.scenes[handle]
	if !ok {
		panic("scene not found: " + handle)
	}

	if g.currentScene == nil {
		g.currentScene = scene
		fmt.Println("Game::SetScene to " + handle)
	} else {
		g.nextScene = scene
		fmt.Println("Game::SetScene updated to " + handle)
	}
}
