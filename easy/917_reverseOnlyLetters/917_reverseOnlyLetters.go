package reverseOnlyLetters

/*
Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.



Example 1:

Input: "ab-cd"
Output: "dc-ba"

Example 2:

Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"

Example 3:

Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"



Note:

    S.length <= 100
    33 <= S[i].ASCIIcode <= 122
    S doesn't contain \ or "


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Only Letters.
// Memory Usage: 2 MB, less than 69.23% of Go online submissions for Reverse Only Letters.

func reverseOnlyLetters(S string) string {
	if S == "" {
		return ""
	}
	moveForward := 0
	moveBack := len(S) - 1
	res := []byte(S)
	for {
		for ; moveForward < len(S)-1; moveForward++ {
			if isLetter(res[moveForward]) {
				break
			}
		}

		for ; moveBack > 0; moveBack-- {
			if isLetter(S[moveBack]) {
				break
			}
		}

		if moveBack == 0 || moveForward == len(S)-1 {
			if isLetter(S[moveBack]) && isLetter(res[moveForward]) {
				res[moveForward] = S[moveBack]
			}
			break
		}

		res[moveForward] = S[moveBack]
		moveForward++
		moveBack--
	}
	return string(res)

}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
