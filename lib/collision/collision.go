package collision

import (
	"fmt"
	"math"

	"github.com/co0p/tankism/lib/ecs"
)

type Point struct {
	X, Y float64
}

type Edges [4]Point

func (e1 *Edges) Equal(e2 Edges) bool {
	same := true

	for k, _ := range e1 {
		a := e1[k]
		b := e2[k]

		if a.X != b.X || a.Y != b.Y {
			same = false
		}
	}

	return same
}

func (e *Edges) String() string {
	return fmt.Sprintf("[[%f, %f], [%f, %f], [%f, %f], [%f, %f]]",
		e[0].X, e[0].Y, e[1].X, e[1].Y, e[2].X, e[2].Y, e[3].X, e[3].Y)
}

type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64

	p1, p2, p3, p4 Point

	E *ecs.Entity // #38 - TODO remove, collision should work with primitives and not know about about ecs*
}

// Rotate sets the edges of the rectangle based on the origin X,Y, width and height and the given rotation in deg by
// multipliying with rotation matrix. See https://ebitengine.org/en/documents/matrix.html
func (b *BoundingBox) Rotate(theta float64) BoundingBox {

	// ebiten has y value pointing down, aka lefthanded
	b.p1 = Point{X: b.X, Y: b.Y}
	b.p2 = Point{X: b.X + b.Width, Y: b.Y}
	b.p3 = Point{X: b.X, Y: b.Y + b.Height}
	b.p4 = Point{X: b.X + b.Width, Y: (b.Y + b.Height)}

	if theta == 0 {
		return *b
	}

	sin, cos := math.Sincos(-theta)
	points := []*Point{&b.p1, &b.p2, &b.p3, &b.p4}
	for _, v := range points {
		/* rotate x,y by deg.
		|cos, -sin| * |X|		=	|cos*X + -sin*Y|
		|sin, cos|    |Y|			|sin*X* + cos*Y|
		*/
		v.X = cos*v.X + -sin*v.Y
		v.Y = sin*v.X + cos*v.Y
	}

	return *b
}

func (b *BoundingBox) Edges() Edges {
	return Edges{b.p1, b.p2, b.p3, b.p4}
}

// TODO: replace with Separating Axis Theorem to accomodate for rotation of rectangles
// TODO: easier to grok: https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect and do this
// for each line of a rectacle vs all lines of the other

func AABBCollision(a, b BoundingBox) bool {
	return a.X < b.X+b.Width &&
		a.X+a.Width > b.X &&
		a.Y < b.Y+b.Height &&
		a.Y+a.Height > b.Y
}
