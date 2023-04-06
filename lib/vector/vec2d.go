package vector

import "math"

type Vec2d struct {
	X, Y float64
}

/*
	 Rotate rotates the vector x,y by theta (rad) clockwise, see https://ebitengine.org/en/documents/matrix.html#Rotating

		|cos,-sin| * |X|	=	|cos*X - sin*Y|
		|sin, cos|   |Y|		|sin*X + cos*Y|
*/
func (v *Vec2d) Rotate(theta float64) Vec2d {
	sin, cos := math.Sincos(theta)
	x := math.Round(cos*v.X - sin*v.Y)
	y := math.Round(sin*v.X + cos*v.Y)

	return Vec2d{x, y}
}

// XY returns the X, Y values as a pair of vector
func (v *Vec2d) XY() (float64, float64) {
	return v.X, v.Y
}

// MidPoint returns the midpoint as a vector of the given two vectors
func MidPoint(v1 Vec2d, v2 Vec2d) Vec2d {
	mx := v1.X + v2.X
	my := v1.Y + v2.Y
	return Vec2d{mx / 2, my / 2}
}
