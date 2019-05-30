package canPlaceFlowers

/*
Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.

Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: True

Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: False

Note:

    The input array won't violate no-adjacent-flowers rule.
    The input array size is in the range of [1, 20000].
    n is a non-negative integer which won't exceed the input array size.

*/

// Runtime: 16 ms, faster than 99.12% of Go online submissions for Can Place Flowers.
// Memory Usage: 5.9 MB, less than 15.79% of Go online submissions for Can Place Flowers.

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		return flowerbed[0] == 0 && n <= 1 || flowerbed[0] == 1 && n == 0
	}
	avail := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i == 0 {
				if flowerbed[i+1] == 0 {
					avail++
					flowerbed[i] = 1
				}
				continue
			}

			if i == len(flowerbed)-1 {
				if flowerbed[i-1] == 0 {
					avail++
					flowerbed[i] = 1
				}
				continue
			}

			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				avail++
				flowerbed[i] = 1
			}
		}
	}

	return avail >= n
}
