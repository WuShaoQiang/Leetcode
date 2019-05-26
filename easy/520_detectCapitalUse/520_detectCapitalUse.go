package detectCapitalUse

/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

    All letters in this word are capitals, like "USA".
    All letters in this word are not capitals, like "leetcode".
    Only the first letter in this word is capital, like "Google".

Otherwise, we define that this word doesn't use capitals in a right way.



Example 1:

Input: "USA"
Output: True



Example 2:

Input: "FlaG"
Output: False



Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Detect Capital.
// Memory Usage: 2.1 MB, less than 27.78% of Go online submissions for Detect Capital.

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	// the first letter is lowercase, so the rest of its letters should be lowercase
	if word[0] >= 'a' && word[0] <= 'z' {
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
		// the first letter is uppercase, judge the second letter
	} else if word[0] >= 'A' && word[0] <= 'Z' {
		// the second letter is uppercase, so the rest of them should be uppercase
		if word[1] >= 'A' && word[1] <= 'Z' {
			for i := 2; i < len(word); i++ {
				if word[i] < 'A' || word[i] > 'Z' {
					return false
				}
			}
		} else {
			for i := 2; i < len(word); i++ {
				if word[i] < 'a' || word[i] > 'z' {
					return false
				}
			}
		}
	}

	return true
}
