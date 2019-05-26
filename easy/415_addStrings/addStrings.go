package addStrings

/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

    The length of both num1 and num2 is < 5100.
    Both num1 and num2 contains only digits 0-9.
    Both num1 and num2 does not contain any leading zero.
    You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Strings.
// Memory Usage: 2.5 MB, less than 83.94% of Go online submissions for Add Strings.

func addStrings(num1 string, num2 string) string {
	// 保持num1长度大于或等于num2长度
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	dis := len(num1) - len(num2)
	res := make([]byte, len(num1)+1)
	idx := len(res) - 1
	carry := false

	for i := len(num2) - 1; i >= 0; i-- {
		tmp := num1[i+dis] - '0' + num2[i] - '0'
		if tmp >= 10 {
			if carry {
				tmp = tmp%10 + 1
			} else {
				tmp = tmp % 10
				carry = true
			}
		} else {
			if carry {
				if tmp == 9 {
					tmp = 0
					carry = true
				} else {
					tmp = tmp + 1
					carry = false
				}
			}
		}
		res[idx] = tmp + '0'
		idx--
	}

	for i := len(num1) - len(num2) - 1; i >= 0; i-- {
		tmp := num1[i] - '0'
		if carry {
			tmp++
		}

		if tmp == 10 {
			tmp = 0
			carry = true
		} else {
			carry = false
		}

		res[idx] = tmp + '0'
		idx--
	}

	if carry {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])

}
