package collision

import (
	"math"

	"github.com/co0p/tankism/lib/ecs"
	"github.com/co0p/tankism/lib/vector"
)

type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64

	p1, p2, p3, p4 vector.Vec2d

	E *ecs.Entity // #38 - TODO remove, collision should work with primitives and not know about about ecs*
}

// Rotate sets the edges of the rectangle based on the origin X,Y, width and height and the given rotation in deg by
// multipliying with rotation matrix. See https://ebitengine.org/en/documents/matrix.html
func (b *BoundingBox) Rotate(theta float64) BoundingBox {

	// set
	b.Edges()

	// rotate
	b.p1 = b.p1.Rotate(theta)
	b.p2 = b.p2.Rotate(theta)
	b.p3 = b.p3.Rotate(theta)
	b.p4 = b.p4.Rotate(theta)

	return *b
}

func (b *BoundingBox) Edges() [4]vector.Vec2d {
	b.p1 = vector.Vec2d{b.X, b.Y}
	b.p2 = vector.Vec2d{b.X + b.Width, b.Y}
	b.p3 = vector.Vec2d{b.X + b.Width, b.Y + b.Height}
	b.p4 = vector.Vec2d{b.X, b.Y + b.Height}

	return [4]vector.Vec2d{b.p1, b.p2, b.p3, b.p4}
}

// MinArea returns a new bounding box covering the original boundingbox based incorperating rotation
// see https://i.stack.imgur.com/ExZl3.png
func (b *BoundingBox) MinArea() BoundingBox {

	b.Edges()

	minX := b.p1.X
	maxX := b.p1.X
	minY := b.p1.Y
	maxY := b.p1.Y
	for _, v := range [4]vector.Vec2d{b.p1, b.p2, b.p3, b.p4} {
		if v.X < minX {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}

	return BoundingBox{
		X:      minX,
		Y:      minY,
		Width:  math.Abs(maxX) - math.Abs(minX),
		Height: math.Abs(maxY) - math.Abs(minY),
	}

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
