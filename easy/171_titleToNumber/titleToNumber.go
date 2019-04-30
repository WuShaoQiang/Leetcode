package titleToNumber

import "fmt"

func titleToNumber(s string) int {
	result := 0
	tmp := 1
	for i := len(s) - 1; i >= 0; i-- {
		num := int(s[i] - 64)
		fmt.Printf("%v : %v\n", string(s[i]), num)
		result += (num * tmp)
		tmp = tmp * 26
	}
	return result
}
