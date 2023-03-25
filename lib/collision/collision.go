package collision

import (
	"github.com/co0p/tankism/lib/ecs"
)

type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64

	p1, p2, p3, p4 Vec2d

	E *ecs.Entity // #38 - TODO remove, collision should work with primitives and not know about about ecs*
}

// Rotate sets the edges of the rectangle based on the origin X,Y, width and height and the given rotation in deg by
// multipliying with rotation matrix. See https://ebitengine.org/en/documents/matrix.html
func (b *BoundingBox) Rotate(theta float64) BoundingBox {

	// set
	b.p1 = Vec2d{b.X, b.Y}
	b.p2 = Vec2d{b.X + b.Width, b.Y}
	b.p3 = Vec2d{b.X + b.Width, b.Y + b.Height}
	b.p4 = Vec2d{b.X, b.Y + b.Height}

	// rotate
	b.p1.Rotate(theta)
	b.p2.Rotate(theta)
	b.p3.Rotate(theta)
	b.p4.Rotate(theta)

	return *b
}

func (b *BoundingBox) Edges() [4]Vec2d {
	return [4]Vec2d{b.p1, b.p2, b.p3, b.p4}
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
