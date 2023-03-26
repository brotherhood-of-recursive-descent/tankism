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

// TODO MinArea should describe the convex hull of the rectangle ignoring rotation
func Test_Collision_MinArea(t *testing.T) {

	// given
	bbA := collision.BoundingBox{X: 10, Y: 10, Width: 10, Height: 10}
	rotated := bbA.Rotate(math.Pi)

	// when
	coverBox := rotated.MinArea()

	// then ... TODO
	if coverBox.X != 10 {
		t.Errorf("%v", coverBox)

	}

}

func givenTwoCollidingRectangles() (collision.BoundingBox, collision.BoundingBox) {
	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 150, Y: 150, Width: 100, Height: 100}
	return bbA, bbB
}
