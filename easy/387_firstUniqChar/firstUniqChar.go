package firstUniqChar

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
