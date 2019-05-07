package reverseVowels

func reverseVowels(s string) string {
	b := []byte(s)
	// m := map[byte]bool{
	// 	'a': true,
	// 	'e': true,
	// 	'i': true,
	// 	'o': true,
	// 	'u': true,
	// 	'A': true,
	// 	'E': true,
	// 	'I': true,
	// 	'O': true,
	// 	'U': true,
	// }

	i := 0
	j := len(b) - 1
	for {
		for ; i < len(b); i++ {
			if isVowel(b[i]) {
				break
			}
		}

		for ; j >= 0; j-- {
			if isVowel(b[j]) {
				break
			}
		}

		if i >= j {
			break
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func isVowel(b byte) bool {

	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'

}
