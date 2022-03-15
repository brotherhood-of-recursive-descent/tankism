package menu

import (
	"fmt"
	"github.com/co0p/tankism"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuScene struct {
	WindowWidth  int
	WindowHeight int

	currentIndex int
	playButton   *tankism.Button
	exitButton   *tankism.Button

	backgroundImage *tankism.MenuImage
}

func NewMenuScene(*tankism.SceneManager) *MenuScene {
	scene := &MenuScene{}

	playAction := func() {
		fmt.Println("play action called")
	}

	exitAction := func() {
		fmt.Println("exit action called")
	}

	scene.backgroundImage = tankism.NewMenuImage(scene)
	scene.playButton = tankism.NewButton("play", 300, 300, playAction)
	scene.exitButton = tankism.NewButton("exit", 300, 500, exitAction)

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
