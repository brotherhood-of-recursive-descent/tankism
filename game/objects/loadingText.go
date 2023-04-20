package objects

import (
	"image/color"

	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// LoadingText is the loading text animation for the start screen.
type LoadingText struct {
	text           string
	font           font.Face
	primaryColor   color.Color
	secondaryColor color.Color
	frameCount     int
}

func NewLoadingText() *LoadingText {

	return &LoadingText{
		text:           "Loading ...",
		font:           resources.FontBig,
		primaryColor:   lib.ColorYellow,
		secondaryColor: nil,
		frameCount:     0,
	}
}

func (l *LoadingText) Draw(screen *ebiten.Image) {
	w, h := lib.WidthHeight(screen)
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
