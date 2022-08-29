package game

import (
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is the shell of a game and is a meant to be used as a container for scenes
type Game struct {
	sceneManager lib.SceneManager
}

func (g *Game) AddScene(handle string, scene lib.Scene) {
	g.sceneManager.RegisterScene(handle, scene)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.CurrentScene().Draw(screen)
}

func (g *Game) Update() error {
	return g.sceneManager.CurrentScene().Update()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	return &Game{
		sceneManager: *lib.NewSceneManager(),
	}
}
