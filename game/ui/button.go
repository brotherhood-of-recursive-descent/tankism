package ui

import (
	"image/color"

	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	text            string
	font            font.Face
	posX, posY      int
	textColor       color.Color
	backgroundColor color.Color
	activeColor     color.Color
	Active          bool

	action func()
}

func NewButton(label string, posX int, posY int, action func()) *Button {

	return &Button{
		text:        label,
		font:        resources.FontHuge,
		textColor:   lib.ColorYellow,
		activeColor: lib.ColorRed,
		posX:        posX,
		posY:        posY,
		Active:      false,
		action:      action,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	if b.Active {
		text.Draw(screen, b.text, b.font, b.posX, b.posY, b.activeColor)
	} else {
		text.Draw(screen, b.text, b.font, b.posX, b.posY, b.textColor)
	}
}

func (b *Button) Update() error {
	return nil
}
