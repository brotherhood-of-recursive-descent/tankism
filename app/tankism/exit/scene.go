package exit

import (
	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ALPHA_MIN  = 0
	ALPHA_STEP = 0.005
)

type ExitScene struct {
	game.GameScene

	game         *game.Game
	currentAlpha float64
	prevImage    *ebiten.Image
}

func NewExitScene(game *game.Game) *ExitScene {
	img, _ := resources.LoadImage(resources.BackgroundImage)

	return &ExitScene{
		game:         game,
		prevImage:    ebiten.NewImageFromImage(img),
		currentAlpha: 1,
	}
}

func (s *ExitScene) Draw(screen *ebiten.Image) {
	if s.prevImage == nil {
		s.prevImage = ebiten.NewImageFromImage(screen)
	}
	op := ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(float32(s.currentAlpha))
	screen.DrawImage(s.prevImage, &op)
}

func (s *ExitScene) Update() error {
	if s.currentAlpha > ALPHA_MIN {
		s.currentAlpha -= ALPHA_STEP
	}
	if s.currentAlpha <= ALPHA_MIN {
		s.game.Exit()
	}
	return nil
}
