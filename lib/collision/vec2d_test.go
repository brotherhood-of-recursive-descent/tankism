package collision_test

import (
	"math"
	"testing"

	"github.com/co0p/tankism/lib/collision"
)

func Test_Vec2d_Rotate(t *testing.T) {

	cases := []struct {
		desc     string
		vec      collision.Vec2d
		theta    float64
		expected collision.Vec2d
	}{
		{"0° rotation",
			collision.Vec2d{1, 1},
			0,
			collision.Vec2d{1, 1},
		},
		{"90° rotation",
			collision.Vec2d{1, 1},
			0.5 * math.Pi,
			collision.Vec2d{-1, 1},
		},
		{"180° rotation",
			collision.Vec2d{1, 1},
			1 * math.Pi,
			collision.Vec2d{1, -1},
		},
		{"270° rotation",
			collision.Vec2d{1, 1},
			1.5 * math.Pi,
			collision.Vec2d{1, -1},
		},
		{"360° rotation",
			collision.Vec2d{1, 1},
			2 * math.Pi,
			collision.Vec2d{1, 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {

			actual := tc.vec.Rotate(tc.theta)

			if actual.X != tc.expected.X || actual.Y != tc.expected.Y {
				t.Errorf("expected: \n%v, got:\n%v\n", tc.expected, actual)
			}
		})
	}
}
