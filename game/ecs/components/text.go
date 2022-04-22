package components

import (
	"image/color"

	"github.com/co0p/tankism/lib/ecs"
	"golang.org/x/image/font"
)

const TextType = "text"

// Text holds all information needed to render text
type Text struct {
	Value string
	Color color.Color
	Font  font.Face
}

func (t Text) Type() ecs.ComponentType {
	return TextType
}
