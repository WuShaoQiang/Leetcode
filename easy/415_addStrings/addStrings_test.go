package addStrings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddStrings(t *testing.T) {
	type Input struct {
		num1 string
		num2 string
	}
	tests := map[string]struct {
		input Input
		want  string
	}{
		"Test 1": {
			input: Input{
				num1: "1",
				num2: "999",
			},
			want: "1000",
		},

		"Test 2": {
			input: Input{
				num1: "111",
				num2: "999",
			},
			want: "1110",
		},

		"Test 3": {
			input: Input{
				num1: "1111",
				num2: "999",
			},
			want: "2110",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := addStrings(tc.input.num1, tc.input.num2)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
