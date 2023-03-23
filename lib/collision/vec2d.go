package collision

import "math"

type Vec2d struct {
	X, Y float64
}

func (v *Vec2d) Rotate(theta float64) Vec2d {
	if int(theta)%360 == 0 {
		return *v
	}

	sin, cos := math.Sincos(theta)
	/* rotate x,y
	| cos, sin| * |X|		=	| cos*X + sin*Y|
	|-sin, cos|   |Y|			|-sin*X + cos*Y|
	*/
	v.X = math.Round(cos*v.X + sin*v.Y)
	v.Y = math.Round(-sin*v.X + cos*v.Y)
	return *v
}
