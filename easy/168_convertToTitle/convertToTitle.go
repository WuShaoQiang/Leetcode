package convertToTitle

import (
	"strings"
)

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}

	len := 0
	tmp := n
	for tmp/26 > 0 {
		len++
		tmp = tmp / 26
	}
	strs := make([]string, len+1)

	for n > 0 {
		n--
		strs[len] = string((n % 26) + 'A')
		n = n / 26
		len--
	}
	return strings.Join(strs, "")

}
