package collision

import "math"

type Vec2d struct {
	X, Y float64
}

/* Rotate rotates the vector x,y by theta (rad) clockwise, see https://ebitengine.org/en/documents/matrix.html#Rotating

	|cos,-sin| * |X|	=	|cos*X - sin*Y|
	|sin, cos|   |Y|		|sin*X + cos*Y|
	*/
func (v *Vec2d) Rotate(theta float64) Vec2d {
	sin, cos := math.Sincos(theta)
	x := math.Round(cos*v.X - sin*v.Y)
	y := math.Round(sin*v.X + cos*v.Y)

	return Vec2d{x,y}
}
