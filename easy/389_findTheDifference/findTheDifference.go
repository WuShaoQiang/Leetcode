package findTheDifference

// func findTheDifference(s string, t string) byte {
// 	m := make([]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		m[s[i]-'a']++
// 		m[t[i]-'a']++
// 	}
// 	m[t[len(t)-1]-'a']++
// 	for idx, count := range m {
// 		if count%2 != 0 {
// 			return byte(idx) + 'a'
// 		}
// 	}
// 	return byte(0)
// }

// func findTheDifference(s string, t string) byte {
// 	sb := []byte(s)
// 	tb := []byte(t)
// 	sSum := 0
// 	tSum := int(tb[len(tb)-1])
// 	for i := 0; i < len(sb); i++ {
// 		sSum += int(sb[i])
// 		tSum += int(tb[i])
// 	}
// 	return byte(tSum - sSum)
// }

func findTheDifference(s string, t string) byte {
	var i int
	result1, result2 := uint8(0), uint8(0)
	for i = 0; i < len(s); i++ {
		result1 ^= s[i]
		result2 ^= t[i]
	}
	result2 ^= t[i]
	return result1 ^ result2
}
