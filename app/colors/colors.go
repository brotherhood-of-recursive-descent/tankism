package main

import (
	"fmt"

	"github.com/co0p/tankism/lib"
)

func main() {

	green := lib.ColorGreen
	r, _, _, _ := green.RGBA()

	fr := float32(float32(r) / 65535.0)
	fR := float32(19018.0 / 65535)

	fmt.Printf("%v, %f, %f \n", r, fr, fR)
}
