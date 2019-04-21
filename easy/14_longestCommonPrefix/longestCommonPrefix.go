package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	}

	minPos := 0
	minLen := len(strs[0])
	for i := 0; i < size; i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
			minPos = i
		}
	}

	for j := 0; j < minLen; j++ {
		temp := strs[minPos][j]
		for i := 0; i < size; i++ {
			if i == minPos {
				continue
			}
			if temp != strs[i][j] {
				if j == 0 {
					return ""
				}
				return strs[i][:j]
			}
		}
	}
	return strs[minPos]

}
