package findRestaurant

import "math"

/*

Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

Example 1:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".
Example 2:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
Note:
The length of both lists will be in the range of [1, 1000].
The length of strings in both lists will be in the range of [1, 30].
The index is starting from 0 to the list length minus 1.
No duplicates in both lists.
*/

// Runtime: 28 ms, faster than 100.00% of Go online submissions for Minimum Index Sum of Two Lists.
// Memory Usage: 6.8 MB, less than 62.50% of Go online submissions for Minimum Index Sum of Two Lists.

func findRestaurant(list1 []string, list2 []string) []string {
	m1 := make(map[string]int)
	for idx, restaurant := range list1 {
		m1[restaurant] = idx
	}
	sum := math.MaxInt64

	res := []string{}

	for idx2, l2 := range list2 {
		if idx1, ok := m1[l2]; ok {
			if sum == idx2+idx1 {
				res = append(res, l2)
			} else if sum > idx1+idx2 {
				sum = idx1 + idx2
				res = []string{l2}
			}
		}
	}
	return res
}
