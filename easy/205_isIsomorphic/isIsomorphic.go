package isIsomorphic

func isIsomorphic(s string, t string) bool {
	beUsed := make(map[byte]bool)
	letterMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if b, exist := letterMap[s[i]]; !exist {
			if _, used := beUsed[t[i]]; used {
				return false
			}
			beUsed[t[i]] = true
			letterMap[s[i]] = t[i]
		} else {
			if b != t[i] {
				return false
			}
		}
	}
	return true
}
