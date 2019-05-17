package addBinary

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"

Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
// Memory Usage: 2.2 MB, less than 84.50% of Go online submissions for Add Binary.

func addBinary(a string, b string) string {
	// 为了下面编程方便,选择b长a短
	if len(a) > len(b) {
		a, b = b, a
	}
	// 差距
	dis := len(b) - len(a)
	res := make([]byte, len(b)+1)
	// 定义res的索引
	idx := len(res) - 1
	var carry bool
	// 基于小的长度遍历
	for i := len(a) - 1; i >= 0; i-- {
		// 有进位的情况
		if carry {
			switch {
			case a[i] == '1' && b[i+dis] == '1':
				res[idx] = '1'
				idx--
				carry = true
			case a[i] == '1' || b[i+dis] == '1':
				res[idx] = '0'
				idx--
				carry = true

			case a[i] == '0' && b[i+dis] == '0':
				res[idx] = '1'
				idx--
				carry = false
			}
			// 没有进位
		} else {
			switch {
			case a[i] == '1' && b[i+dis] == '1':
				res[idx] = '0'
				idx--
				carry = true

			case a[i] == '1' || b[i+dis] == '1':
				res[idx] = '1'
				idx--
				carry = false

			case a[i] == '0' && b[i+dis] == '0':
				res[idx] = '0'
				idx--
				carry = false
			}
		}
	}

	// 开始对长的字符串单独操作
	for i := len(b) - len(a) - 1; i >= 0; i-- {
		switch {
		case b[i] == '0' && carry:
			res[idx] = '1'
			idx--
			carry = false
		case b[i] == '1' && carry:
			res[idx] = '0'
			idx--
			carry = true
		case b[i] == '0' && !carry:
			res[idx] = '0'
			idx--
			carry = false
		case b[i] == '1' && !carry:
			res[idx] = '1'
			idx--
			carry = false
		}
	}

	if carry {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])

}
