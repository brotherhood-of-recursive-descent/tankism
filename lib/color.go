package lib

import (
	"image/color"
)

var ColorBlack = Color{R: 0, G: 0, B: 0, A: 0}
var ColorWhite = Color{R: 255, G: 255, B: 255, A: 255}
var ColorYellow = Color{R: 0xfa, G: 0xfa, B: 0xd2, A: 0xff}
var ColorRed = Color{R: 250, G: 37, B: 37, A: 255}
var ColorGreen = Color{R: 0x4a, G: 0xf6, B: 0x26, A: 0xff}
var ColorBlue = Color{R: 37, G: 37, B: 255, A: 0xff}

// reimplementing Color from color.Color to make it json deserializeable ... dunno why needed
type Color struct {
	R, G, B, A uint8
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

func GetRGBA64(clr color.Color) (r1, g1, b1, a1 float64) {
	r, g, b, a := clr.RGBA()
	return float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0
}
