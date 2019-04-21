package romanToInt

func romanToInt(s string) int {
	var sum int
	mp := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := 0; i < len(s)-1; i++ {
		if mp[string(s[i])] < mp[string(s[i+1])] {
			sum -= mp[string(s[i])]
		} else {
			sum += mp[string(s[i])]
		}
	}
	return sum + mp[string(s[len(s)-1])]
}
