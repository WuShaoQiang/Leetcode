package convertToTitle

import (
	"strings"
)

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

Example 1:

Input: 1
Output: "A"

Example 2:

Input: 28
Output: "AB"

Example 3:

Input: 701
Output: "ZY"
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Excel Sheet Column Title.
// Memory Usage: 1.9 MB, less than 43.27% of Go online submissions for Excel Sheet Column Title.

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
