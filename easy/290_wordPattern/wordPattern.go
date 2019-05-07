package wordPattern

import "strings"

// 要注意是双边映射
func wordPattern(pattern string, str string) bool {
	patterns := strings.Split(pattern, "")
	strs := strings.Split(str, " ")
	if len(patterns) != len(strs) {
		return false
	}
	pattrenMap := make(map[string]string)
	beUsed := make(map[string]bool)
	for i := 0; i < len(patterns); i++ {
		if _, exist := pattrenMap[patterns[i]]; !exist {
			if _, isUsed := beUsed[strs[i]]; isUsed {
				return false
			}
			pattrenMap[patterns[i]] = strs[i]
			beUsed[strs[i]] = true

		} else {
			if pattrenMap[patterns[i]] != strs[i] {
				return false
			}
		}
	}
	return true

}
