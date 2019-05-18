package firstUniqChar

/*
 Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.

Note: You may assume the string contain only lowercase letters. */

// Runtime: 8 ms, faster than 98.65% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.7 MB, less than 35.77% of Go online submissions for First Unique Character in a String.

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, b := range s {
		m[b-'a']++
	}

	for idx, b := range s {
		if m[b-'a'] == 1 {
			return idx
		}
	}

	return -1
}

// func firstUniqChar(s string) int {
//     for idx,r := range s{
//         if strings.Index(s,string(r)) == strings.LastIndex(s,string(r)){
//             return idx
//         }
//     }

//     return -1
// }
