package collision_test

import (
	"math"
	"testing"

	"github.com/co0p/tankism/lib/collision"
)

func Test_CollisionAABB_collision(t *testing.T) {

	bbA, bbB := givenTwoCollidingRectangles()

	ok := collision.AABBCollision(bbA, bbB)
	if !ok {
		t.Errorf("expected collision between bbA and bbB, got false\n")
	}
}

func Test_CollisionAABB_no_collision(t *testing.T) {

	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 400, Y: 400, Width: 100, Height: 100}

	ok := collision.AABBCollision(bbA, bbB)
	if ok {
		t.Errorf("expected no collision between bbA and bbB, got true\n")
	}
}

func Test_Collision_MinArea_Same(t *testing.T) {

	// given
	bbA := collision.BoundingBox{X: 10, Y: 10, Width: 10, Height: 10}

	// when
	coverBox := bbA.MinArea()

	// then ... TODO
	if coverBox.X != bbA.X || coverBox.Y != bbA.Y ||
		coverBox.Width != bbA.Width || coverBox.Height != bbA.Height {
		t.Errorf("expected minArea to be %v, got %v\n", bbA, coverBox)
	}

}

// TODO MinArea should describe the convex hull of the rectangle ignoring rotation
func Test_Collision_MinArea(t *testing.T) {

	// given
	bbA := collision.BoundingBox{X: 10, Y: 20, Width: 20, Height: 10}
	rotated := bbA.Rotate(1 / 4 * math.Pi) // 45°

	// when
	coverBox := rotated.MinArea()
	edges := coverBox.Edges()

	// then ... TODO
	if edges[0].X != 1 {
		t.Errorf("%v", edges)
	}
}

func givenTwoCollidingRectangles() (collision.BoundingBox, collision.BoundingBox) {
	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 150, Y: 150, Width: 100, Height: 100}
	return bbA, bbB
}
