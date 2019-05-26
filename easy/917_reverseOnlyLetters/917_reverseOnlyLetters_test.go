package reverseOnlyLetters

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseOnlyLetters(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Test 1": {
			input: "ab-cd",
			want:  "dc-ba",
		},

		"Test 2": {
			input: "a-bC-dEf-ghIj",
			want:  "j-Ih-gfE-dCba",
		},

		"Test 3": {
			input: "Test1ng-Leet=code-Q!",
			want:  "Qedo1ct-eeLg=ntse-T!",
		},

		"Test 4": {
			input: "",
			want:  "",
		},

		"Test 5": {
			input: "-",
			want:  "-",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverseOnlyLetters(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
