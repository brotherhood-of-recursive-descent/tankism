package menu

import (
	"fmt"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/objects"
	"github.com/co0p/tankism/game/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuScene struct {
	game.GameScene

	game *game.Game

	currentIndex    int
	playButton      *ui.Button
	exitButton      *ui.Button
	backgroundImage *objects.MenuImage
}

func NewMenuScene(game *game.Game) *MenuScene {

	playAction := func() {
		fmt.Println("play action called")
		game.SetScene("SINGLEPLAYER")
	}
	exitAction := func() {
		fmt.Println("exit action called")
		game.SetScene("EXIT")
	}

	return &MenuScene{
		game:            game,
		backgroundImage: objects.NewMenuImage(),
		playButton:      ui.NewButton("play", 300, 300, playAction),
		exitButton:      ui.NewButton("exit", 300, 500, exitAction),
	}
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	m.backgroundImage.Draw(screen)

	m.playButton.Draw(screen)
	m.exitButton.Draw(screen)
}

func (m *MenuScene) HandleInput() {

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.currentIndex++
		if m.currentIndex > 1 {
			m.currentIndex = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.currentIndex--
		if m.currentIndex < 1 {
			m.currentIndex = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		fmt.Println("execute action")
		if m.currentIndex == 0 {
			m.game.SetScene("SINGLEPLAYER")
		}
		if m.currentIndex == 1 {
			m.game.SetScene("EXIT")
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
}
