package start

import (
	"fmt"
	"github.com/co0p/tankism/game/objects"
	"github.com/co0p/tankism/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type StartScene struct {
	WindowWidth  int
	WindowHeight int
	startTime    time.Time

	sceneManager *lib.SceneManager
	nextScene    lib.Scene

	loadingImage *objects.LoadingImage
	loadingText  *objects.LoadingText
}

func NewStartScreen(sceneManager *lib.SceneManager, nextScene lib.Scene) *StartScene {

	scene := StartScene{}
	scene.sceneManager = sceneManager
	scene.nextScene = nextScene

	scene.loadingImage = objects.NewLoadingImage(&scene)
	scene.loadingText = objects.NewLoadingText(&scene)

	return &scene
}

func (s *StartScene) Init() error {
	s.startTime = time.Now()
	return nil
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	s.loadingImage.Draw(screen)
	s.loadingText.Draw(screen)
}

func (s *StartScene) Update() error {
	err := s.loadingImage.Update()
	err = s.loadingText.Update()

	if s.loadingDone() {
		fmt.Println("Loading done")
		s.sceneManager.ChangeScene(s.nextScene)
	}

	return err
}

func (s *StartScene) WindowDimension() (int, int) {
	return s.WindowWidth, s.WindowHeight
}

func (s *StartScene) SetWindowDimension(w, h int) {
	s.WindowWidth = w
	s.WindowHeight = h
}

func (s *StartScene) loadingDone() bool {
	startPlus5 := s.startTime.Add(time.Second * 5)
	return time.Now().After(startPlus5)
}
