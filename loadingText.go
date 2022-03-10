package tankism

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

var colorYellow = color.RGBA{R: 0xfa, G: 0xfa, B: 0xd2, A: 0xff}

// LoadingText is the loading text animation for the start screen.
type LoadingText struct {
	scene          Scene
	text           string
	font           font.Face
	primaryColor   color.Color
	secondaryColor color.Color
	frameCount     int
}

func NewLoadingText(scene Scene) *LoadingText {

	return &LoadingText{
		scene:          scene,
		text:           "Loading ...",
		font:           FontBig,
		primaryColor:   colorYellow,
		secondaryColor: nil,
		frameCount:     0,
	}
}

func (l *LoadingText) Draw(screen *ebiten.Image) {
	w, h := l.scene.WindowDimension()
	text.Draw(screen, l.text, l.font, w-300, h-100, l.primaryColor)
}

func (l *LoadingText) Update() error {
	l.frameCount++

	if l.frameCount == 20 {
		l.text = "Loading ."
	}

	if l.frameCount == 40 {
		l.text = "Loading .."
	}

	if l.frameCount == 60 {
		l.text = "Loading ..."
	}

	if l.frameCount == 80 {
		l.frameCount = 0
	}
	return nil
}
