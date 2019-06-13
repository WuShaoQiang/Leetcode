package numUniqueEmails

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumUniqueEmails(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		// "Test 1": {
		// 	input: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
		// 	want:  2,
		// },

		"Test 2": {
			input: []string{"testemail@leetcode.com", "testemail1@leetcode.com", "testemail+david@lee.tcode.com"},
			want:  3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numUniqueEmails(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
