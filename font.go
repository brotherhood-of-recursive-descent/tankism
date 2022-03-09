package tankism

import (
	_ "embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/fonts/test.ttf
var gameFont []byte

var (
	FontSmall  = loadFont(gameFont, 14)
	FontMedium = loadFont(gameFont, 24)
	FontBig    = loadFont(gameFont, 46)
	FontHuge   = loadFont(gameFont, 96)
)

func loadFont(ttfFont []byte, size float64) font.Face {

	tt, err := opentype.Parse(ttfFont)
	if err != nil {
		log.Fatal(err)
	}

	f, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return f
}
