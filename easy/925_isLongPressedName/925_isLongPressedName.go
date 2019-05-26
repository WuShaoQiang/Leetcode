package isLongPressedName

/*
Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.



Example 1:

Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

Example 2:

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

Example 3:

Input: name = "leelee", typed = "lleeelee"
Output: true

Example 4:

Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.



Note:

    name.length <= 1000
    typed.length <= 1000
    The characters of name and typed are lowercase letters.

*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Long Pressed Name.
// Memory Usage: 2 MB, less than 61.11% of Go online submissions for Long Pressed Name.

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}

	if name == typed {
		return true
	}

	i, j := 0, 0
	for i < len(name) {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			j++
		}

		if j == len(typed) && i != len(name) {
			return false
		}
	}

	return true
}
