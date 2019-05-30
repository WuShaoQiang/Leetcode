package canPlaceFlowers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCanPlaceFlowers(t *testing.T) {
	type Input struct {
		flowerbed []int
		n         int
	}
	tests := map[string]struct {
		input Input
		want  bool
	}{
		"Test 1": {
			input: Input{flowerbed: []int{1, 0, 0, 0, 1}, n: 1},
			want:  true,
		},

		"Test 2": {
			input: Input{flowerbed: []int{1, 0, 0, 0, 1}, n: 2},
			want:  false,
		},

		"Test 3": {
			input: Input{flowerbed: []int{1, 0, 0, 0, 0, 1}, n: 2},
			want:  false,
		},

		"Test 4": {
			input: Input{flowerbed: []int{0, 0, 1, 0, 0}, n: 2},
			want:  true,
		},

		"Test 5": {
			input: Input{flowerbed: []int{0}, n: 1},
			want:  true,
		},

		"Test 6": {
			input: Input{flowerbed: []int{0, 1, 0}, n: 1},
			want:  false,
		},

		"Test 7": {
			input: Input{flowerbed: []int{1}, n: 0},
			want:  true,
		},

		"Test 8": {
			input: Input{flowerbed: []int{0}, n: 0},
			want:  true,
		},

		"Test 9": {
			input: Input{flowerbed: []int{1}, n: 1},
			want:  false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := canPlaceFlowers(tc.input.flowerbed, tc.input.n)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
