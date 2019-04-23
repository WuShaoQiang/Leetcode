package addBinary

import (
	"strconv"
)

// func addBinary(a string, b string) string {
// 	if a == "0" {
// 		return b
// 	}

// 	if b == "0" {
// 		return a
// 	}

// 	aNum, _ := strconv.Atoi(a)
// 	bNum, _ := strconv.Atoi(b)

// 	strs := strings.Split(strconv.Itoa(aNum+bNum), "")

// 	add := false

// 	for i := len(strs) - 1; i > 0; i-- {
// 		switch strs[i] {
// 		case "0":
// 			if add {
// 				strs[i] = "1"
// 				add = false
// 			}
// 		case "1":
// 			if add {
// 				strs[i] = "0"
// 				add = true
// 			}
// 		case "2":
// 			if add {
// 				strs[i] = "1"
// 			} else {
// 				strs[i] = "0"
// 			}
// 			add = true
// 		}
// 	}

// 	switch strs[0] {
// 	case "0":
// 		if add {
// 			strs[0] = "1"
// 		}
// 	case "1":
// 		if add {
// 			strs[0] = "10"
// 		}
// 	case "2":
// 		if add {
// 			strs[0] = "11"
// 		} else {
// 			strs[0] = "10"
// 		}
// 	}

// 	return strings.Join(strs, "")

// }

func addBinary(a string, b string) string {

	lengthLonger, lengthShorter := 0, 0
	tmp := ""

	if len(a) > len(b) {
		lengthLonger = len(a)
		lengthShorter = len(b)
	} else {
		lengthLonger = len(b)
		lengthShorter = len(a)
		tmp = b
		b = a
		a = tmp
	}

	resultArray := make([]int, lengthLonger+1)
	carry := 0
	for i := lengthLonger; i >= 0; i-- {
		if i > (lengthLonger - lengthShorter) {
			currentA, _ := strconv.Atoi(a[i-1 : i])
			currentB, _ := strconv.Atoi(b[lengthShorter-lengthLonger+i-1 : lengthShorter-lengthLonger+i])
			resultArray[i] = currentA + currentB + carry
			if resultArray[i] == 2 {
				resultArray[i] = 0
				carry = 1
			} else if resultArray[i] == 3 {
				resultArray[i] = 1
				carry = 1
			} else {
				carry = 0
			}
		} else if i > 0 {
			currentA, _ := strconv.Atoi(a[i-1 : i])
			resultArray[i] = currentA + carry
			if resultArray[i] == 2 {
				resultArray[i] = 0
				carry = 1
			} else {
				carry = 0
			}
		} else {
			resultArray[i] = carry
		}

	}
	resultStr := ""
	for j := 0; j <= lengthLonger; j++ {
		resultStr = resultStr + strconv.Itoa(resultArray[j])
	}
	if resultStr[0:1] == "0" {
		return resultStr[1 : lengthLonger+1]
	} else {
		return resultStr
	}

}
