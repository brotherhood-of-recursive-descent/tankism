package lib

import (
	"image/color"
)

var ColorWhite = color.White
var ColorBlack = color.Black
var ColorYellow = color.RGBA{R: 0xfa, G: 0xfa, B: 0xd2, A: 0xff}
var ColorRed = color.RGBA{R: 0xff, G: 0x4a, B: 0x4a, A: 0xff}
var ColorGreen = color.RGBA{R: 0x4a, G: 0xf6, B: 0x26, A: 0xff}
var ColorBlue = color.RGBA{R: 0x22, G: 0x22, B: 0x4a, A: 0xff}

func GetRGBA64(clr color.Color) (r1, g1, b1, a1 float64) {
	r, g, b, a := clr.RGBA()
	return float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0
}
