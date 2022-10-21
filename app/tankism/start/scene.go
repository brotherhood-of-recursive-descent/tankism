package start

import (
	"time"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

type StartScene struct {
	game.GameScene

	game      *game.Game
	startTime time.Time

	loadingImage *objects.LoadingImage
	loadingText  *objects.LoadingText
}

func NewStartScreen(game *game.Game) *StartScene {
	return &StartScene{
		game:         game,
		loadingImage: objects.NewLoadingImage(),
		loadingText:  objects.NewLoadingText(),
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	s.loadingImage.Draw(screen)
	s.loadingText.Draw(screen)
}

func (s *StartScene) Update() error {

	if s.startTime.IsZero() {
		s.startTime = time.Now()
	}

	s.loadingImage.Update()
	s.loadingText.Update()

	if s.loadingDone() {
		s.game.SetScene("MENU")
	}
	return nil
}

func HandleInput() {}

func (s *StartScene) loadingDone() bool {
	startPlus5 := s.startTime.Add(time.Second * 5)
	return time.Now().After(startPlus5)
}
