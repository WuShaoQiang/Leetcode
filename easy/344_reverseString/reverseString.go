package reverseString

func reverseString(s []byte) {
	tmp := byte(0)
	j := len(s) - 1
	for i := 0; i < j; {
		// if s[i] == s[l-i-1] {
		// 	continue
		// }

		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}

}
