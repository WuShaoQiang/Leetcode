package addDigits

// 迭代
// func addDigits(num int) int {
// 	if num < 10 {
// 		return num
// 	}
// 	res := 0
// 	for res == 0 {
// 		for num > 0 {
// 			res += num % 10
// 			fmt.Println("res:", res)
// 			num = num / 10
// 		}
// 		if res >= 10 {
// 			num = res
// 			res = 0
// 		}
// 	}
// 	return res
// }

// f(10*x+y) = f(9*x+x+y) 所以当我们取余9的时候，就成了x+y
// 如果x+y也正好是9的倍数，则返回9

func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	if num%9 == 0 {
		return 9
	}

	return num % 9
}
