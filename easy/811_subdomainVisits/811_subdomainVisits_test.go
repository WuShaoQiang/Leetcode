package subdomainVisits

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubdomainVisits(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"Test 1": {
			input: []string{"9001 discuss.leetcode.com"},
			want:  []string{"9001 com", "9001 leetcode.com", "9001 discuss.leetcode.com"},
		},

		"Test 2": {
			input: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			want:  []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := subdomainVisits(tc.input)
			diff := cmp.Diff(got, tc.want)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
