package lib

import (
	"image/color"
)

var ColorWhite = color.White
var ColorBlack = color.Black
var ColorYellow = color.RGBA{R: 0xfa, G: 0xfa, B: 0xd2, A: 0xff}
var ColorRed = color.RGBA{R: 250, G: 37, B: 37, A: 255}
var ColorGreen = color.RGBA{R: 0x4a, G: 0xf6, B: 0x26, A: 0xff}
var ColorBlue = color.RGBA{R: 37, G: 37, B: 255, A: 0xff}

func GetRGBA64(clr color.Color) (r1, g1, b1, a1 float64) {
	r, g, b, a := clr.RGBA()
	return float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0
}
