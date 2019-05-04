package isAnagram

func isAnagram(s string, t string) bool {
	x := make([]int, 26)
	y := make([]int, 26)

	for i := 0; i < len(s); i++ {
		x[int(s[i]-'a')]++
	}

	for i := 0; i < len(t); i++ {
		y[int(t[i]-'a')]++
	}

	for i := 0; i < 26; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
