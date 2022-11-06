package components

import (
	"github.com/co0p/tankism/lib"
	"github.com/co0p/tankism/lib/ecs"
	"golang.org/x/image/font"
)

const TextType = "text"

// Text holds all information needed to render text
type Text struct {
	Value string
	Color lib.Color
	Font  font.Face
}

func (t Text) Type() ecs.ComponentType {
	return TextType
}
