package menu

import (
	"fmt"

	"github.com/co0p/tankism/game/objects"
	"github.com/co0p/tankism/game/ui"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuScene struct {
	WindowWidth  int
	WindowHeight int

	currentIndex    int
	playButton      *ui.Button
	exitButton      *ui.Button
	sceneManager    *lib.SceneManager
	backgroundImage *objects.MenuImage
}

func NewMenuScene(sm *lib.SceneManager) *MenuScene {
	scene := &MenuScene{}
	scene.sceneManager = sm

	playAction := func() {
		fmt.Println("play action called")
	}

	exitAction := func() {
		fmt.Println("exit action called")
		sm.ChangeScene("EXIT")
	}

	scene.backgroundImage = objects.NewMenuImage(scene)
	scene.playButton = ui.NewButton("play", 300, 300, playAction)
	scene.exitButton = ui.NewButton("exit", 300, 500, exitAction)

	return scene
}

func (m *MenuScene) Init() error {
	fmt.Println("menu scene init")
	return nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	m.backgroundImage.Draw(screen)

	m.playButton.Draw(screen)
	m.exitButton.Draw(screen)
}

func (m *MenuScene) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		fmt.Println("DOWN")
		m.currentIndex++
		if m.currentIndex > 1 {
			m.currentIndex = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		fmt.Println("UP")
		m.currentIndex--
		if m.currentIndex < 1 {
			m.currentIndex = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		fmt.Println("execute action")
		if m.currentIndex == 1 {
			m.sceneManager.ChangeScene("EXIT")
		}
	}

	switch m.currentIndex {
	case 0:
		{
			m.playButton.Active = true
			m.exitButton.Active = false
			break
		}
	case 1:
		{
			m.playButton.Active = false
			m.exitButton.Active = true
			break
		}
	}

	return nil
}

func (m *MenuScene) WindowDimension() (int, int) {
	return m.WindowWidth, m.WindowHeight
}

func (m *MenuScene) SetWindowDimension(w, h int) {
	m.WindowWidth = w
	m.WindowHeight = h
}
