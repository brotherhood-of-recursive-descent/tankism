package vector_test

import (
	"math"
	"testing"

	"github.com/co0p/tankism/lib/vector"
)

func Test_Vec2d_Rotate(t *testing.T) {

	cases := []struct {
		desc     string
		vec      vector.Vec2d
		theta    float64
		expected vector.Vec2d
	}{
		{"0° rotation",
			vector.Vec2d{1, 1},
			0,
			vector.Vec2d{1, 1},
		},
		{"90° rotation",
			vector.Vec2d{1, 1},
			0.5 * math.Pi,
			vector.Vec2d{-1, 1},
		},
		{"180° rotation",
			vector.Vec2d{1, 1},
			1 * math.Pi,
			vector.Vec2d{-1, -1},
		},
		{"270° rotation",
			vector.Vec2d{1, 1},
			1.5 * math.Pi,
			vector.Vec2d{1, -1},
		},
		{"360° rotation",
			vector.Vec2d{1, 1},
			2 * math.Pi,
			vector.Vec2d{1, 1},
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

func Test_Vec2d_MidPoint(t *testing.T) {

	v1 := vector.Vec2d{1, 1}
	v2 := vector.Vec2d{3, 3}

	expected := vector.Vec2d{2, 2}
	actual := vector.MidPoint(v1, v2)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("expected midpoint of %v, %v to be %v, got %v", v1, v2, expected, actual)
	}
}
