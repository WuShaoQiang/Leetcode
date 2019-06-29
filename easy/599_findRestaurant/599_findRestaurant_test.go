package findRestaurant

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindRestaurant(t *testing.T) {
	type Input struct {
		list1 []string
		list2 []string
	}
	tests := map[string]struct {
		input Input
		want  []string
	}{
		"Test 1": {
			input: Input{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}},
			want: []string{"Shogun"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findRestaurant(tc.input.list1, tc.input.list2)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
