package lengthOfLastWord

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	if !strings.Contains(s, " ") {
		return len(s)
	}

	s = strings.TrimSpace(s)

	lastSpaceIndex := strings.LastIndex(s, " ")
	return len(s) - lastSpaceIndex - 1
}
