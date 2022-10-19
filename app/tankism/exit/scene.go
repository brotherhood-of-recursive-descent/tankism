package exit

import (
	"fmt"
	"os"

	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ALPHA_MIN  = 0
	ALPHA_STEP = 0.005
)

type ExitScene struct {
	WindowWidth  int
	WindowHeight int

	sceneManager *lib.SceneManager
	currentAlpha float64
	prevImage    *ebiten.Image
}

func NewExitScene(sceneManager *lib.SceneManager) *ExitScene {

	scene := ExitScene{}
	scene.sceneManager = sceneManager
	scene.currentAlpha = 1
	return &scene
}

func (s *ExitScene) Init(*lib.SceneManager) error {
	fmt.Println("Loaded exit scene...")
	img, _ := resources.LoadImage(resources.BackgroundImage)
	s.prevImage = ebiten.NewImageFromImage(img)

	return nil
}

func (s *ExitScene) Draw(screen *ebiten.Image) {
	if s.prevImage == nil {
		s.prevImage = ebiten.NewImageFromImage(screen)
	}
	op := &ebiten.DrawImageOptions{}
	op.ColorM.ChangeHSV(1, 1, s.currentAlpha)
	screen.DrawImage(s.prevImage, op)
}

func (s *ExitScene) Update() error {
	if s.currentAlpha > ALPHA_MIN {
		s.currentAlpha -= ALPHA_STEP
	}
	if s.currentAlpha <= ALPHA_MIN {
		os.Exit(0)
	}
	return nil
}

func (s *ExitScene) WindowDimension() (int, int) {
	return s.WindowWidth, s.WindowHeight
}

func (s *ExitScene) SetWindowDimension(w, h int) {
	s.WindowWidth = w
	s.WindowHeight = h
}
