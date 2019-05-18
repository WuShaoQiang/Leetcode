package titleToNumber

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

Example 1:

Input: "A"
Output: 1

Example 2:

Input: "AB"
Output: 28

Example 3:

Input: "ZY"
Output: 701
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Excel Sheet Column Number.
// Memory Usage: 2.2 MB, less than 47.68% of Go online submissions for Excel Sheet Column Number.

func titleToNumber(s string) int {
	result := 0
	tmp := 1
	for i := len(s) - 1; i >= 0; i-- {
		num := int(s[i] - 64)
		// fmt.Printf("%v : %v\n", string(s[i]), num)
		result += (num * tmp)
		tmp = tmp * 26
	}
	return result
}
