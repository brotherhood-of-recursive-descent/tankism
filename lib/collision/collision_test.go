package collision_test

import (
	"math"
	"testing"

	"github.com/co0p/tankism/lib/collision"
)

func Test_CollisionAABB_collision(t *testing.T) {

	bbA := collision.BoundingBox{X: 100, Y: 100, Width: 100, Height: 100}
	bbB := collision.BoundingBox{X: 150, Y: 150, Width: 100, Height: 100}

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

func Test_Rotate(t *testing.T) {

	cases := []struct {
		desc     string
		bb       collision.BoundingBox
		theta    float64
		expected collision.Edges
	}{
		{"0° rotation",
			collision.BoundingBox{X: 0, Y: 0, Width: 40, Height: 20},
			0,
			collision.Edges{{0, 0}, {40, 0}, {40, 20}, {0, 20}}},
		{"90° rotation",
			collision.BoundingBox{X: 0, Y: 0, Width: 40, Height: 20},
			0.5 * math.Pi,
			collision.Edges{{0, 0}, {0, 40}, {-20, 40}, {-20, 0}}},
		{"180° rotation",
			collision.BoundingBox{X: 0, Y: 0, Width: 40, Height: 20},
			1 * math.Pi,
			collision.Edges{{0, 0}, {-40, 0}, {-40, -20}, {0, -20}}},
		{"270° rotation",
			collision.BoundingBox{X: 0, Y: 0, Width: 40, Height: 20},
			1.5 * math.Pi,
			collision.Edges{{0, 0}, {0, -40}, {20, -40}, {20, 0}}},
		{"360° rotation",
			collision.BoundingBox{X: 0, Y: 0, Width: 40, Height: 20},
			2 * math.Pi,
			collision.Edges{{0, 0}, {40, 0}, {40, 20}, {0, 20}}},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {

			actual := tc.bb.Rotate(tc.theta)
			actualEdges := actual.Edges()

			if !actualEdges.Equal(tc.expected) {
				t.Errorf("expected: \n%v, got:\n%v\n", tc.expected, actualEdges)
			}
		})
	}
}
