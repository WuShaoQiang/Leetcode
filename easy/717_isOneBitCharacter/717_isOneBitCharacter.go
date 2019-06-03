package isOneBitCharacter

/*
We have two special characters. The first character can be represented by one bit 0. The second character can be represented by two bits (10 or 11).

Now given a string represented by several bits. Return whether the last character must be a one-bit character or not. The given string will always end with a zero.

Example 1:

Input:
bits = [1, 0, 0]
Output: True
Explanation:
The only way to decode it is two-bit character and one-bit character. So the last character is one-bit character.

Example 2:

Input:
bits = [1, 1, 1, 0]
Output: False
Explanation:
The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.

Note:
1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for 1-bit and 2-bit Characters.
// Memory Usage: 2.8 MB, less than 20.00% of Go online submissions for 1-bit and 2-bit Characters.

func isOneBitCharacter(bits []int) bool {
	return isOneBitChar(bits, 0)
}

func isOneBitChar(bits []int, index int) bool {
	if index >= len(bits) {
		return false
	}
	if bits[index] == 0 {
		if index == len(bits)-1 {
			return true
		}
		return isOneBitChar(bits, index+1)
	}
	return isOneBitChar(bits, index+2)

}
