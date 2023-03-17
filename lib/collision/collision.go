package collision

import "github.com/co0p/tankism/lib/ecs"

type BoundingBox struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	E      *ecs.Entity // #38 - TODO remove, collision should work with primitives and not know about about ecs*
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
