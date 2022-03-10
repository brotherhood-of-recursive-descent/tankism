package start

import (
	"github.com/co0p/tankism"
	"github.com/hajimehoshi/ebiten/v2"
)

type StartScene struct {
	WindowWidth  int
	WindowHeight int

	loadingImage *tankism.LoadingImage
	loadingText  *tankism.LoadingText
}

func NewStartScreen() *StartScene {

	scene := StartScene{}
	scene.loadingImage = tankism.NewLoadingImage(&scene)
	scene.loadingText = tankism.NewLoadingText(&scene)

	return &scene
}

func (s *StartScene) Init() error {
	return nil
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	s.loadingImage.Draw(screen)
	s.loadingText.Draw(screen)
}

func (s *StartScene) Update() error {
	err := s.loadingImage.Update()
	err = s.loadingText.Update()

	return err
}

func (s *StartScene) WindowDimension() (int, int) {
	return s.WindowWidth, s.WindowHeight
}

func (s *StartScene) SetWindowDimension(w, h int) {
	s.WindowWidth = w
	s.WindowHeight = h
}
