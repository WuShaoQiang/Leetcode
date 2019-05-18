package reverseString

/*
Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.

 

Example 1:

Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:

Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

*/

// Runtime: 628 ms, faster than 100.00% of Go online submissions for Reverse String.
// Memory Usage: 8.7 MB, less than 39.79% of Go online submissions for Reverse String.

func reverseString(s []byte) {
	tmp := byte(0)
	j := len(s) - 1
	for i := 0; i < j; {
		// if s[i] == s[l-i-1] {
		// 	continue
		// }

		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}

}
