package collision_test

import (
	"testing"

	"github.com/co0p/tankism/lib/collision"
)

func TestCollisionAABB_collision(t *testing.T) {

	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 150, Y: 150, Width: 100, Height: 100}

	ok := collision.AABBCollision(bbA, bbB)
	if !ok {
		t.Errorf("expected collision between bbA and bbB, got false\n")
	}
}

func TestCollisionAABB_no_collision(t *testing.T) {

	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 400, Y: 400, Width: 100, Height: 100}

	ok := collision.AABBCollision(bbA, bbB)
	if ok {
		t.Errorf("expected no collision between bbA and bbB, got true\n")
	}
}
