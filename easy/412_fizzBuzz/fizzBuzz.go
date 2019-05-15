package fizzBuzz

import (
	"strconv"
)

// 写一个程序，输出从 1 到 n 数字的字符串表示。

// 1. 如果 n 是3的倍数，输出“Fizz”；

// 2. 如果 n 是5的倍数，输出“Buzz”；

// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

// 示例：

// n = 15,

// 返回:
// [
//     "1",
//     "2",
//     "Fizz",
//     "4",
//     "Buzz",
//     "Fizz",
//     "7",
//     "8",
//     "Fizz",
//     "Buzz",
//     "11",
//     "Fizz",
//     "13",
//     "14",
//     "FizzBuzz"
// ]

// 执行用时 : 160 ms, 在Fizz Buzz的Go提交中击败了99.21% 的用户
// 内存消耗 : 69.7 MB, 在Fizz Buzz的Go提交中击败了18.18% 的用户

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			res[i] = "FizzBuzz"
		} else if (i+1)%5 == 0 {
			res[i] = "Buzz"
		} else if (i+1)%3 == 0 {
			res[i] = "Fizz"
		} else {

			res[i] = strconv.Itoa(i + 1)
		}
	}
	return res
}
