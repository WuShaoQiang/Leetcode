package canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	// m := make(map[rune]int)
	m := make([]int, 26)
	for _, b := range magazine {
		m[b-'a']++
	}

	for _, r := range ransomNote {
		if count := m[r-'a']; count > 0 {
			m[r-'a']--
		} else {
			return false
		}
	}
	return true
}
