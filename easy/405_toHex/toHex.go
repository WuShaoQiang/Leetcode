package toHex

func toHex(num int) string {
	m := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := make([]byte, 8)
	for i := 0; i < 8; i++ {
		res[7-i] = m[num&15]
		num = num >> 4
	}
	for i := 0; i < 8; i++ {
		if res[i] != '0' {
			return string(res[i:])
		}
	}
	return "0"
}
