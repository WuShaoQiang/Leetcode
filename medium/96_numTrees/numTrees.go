package numTrees

/*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

func numTrees(n int) int {
	return genTrees(1, n)
}

func genTrees(s, e int) int {
	var sum int

	if s > e {
		return 1
	}

	for i := s; i <= e; i++ {
		left := genTrees(s, i-1)
		right := genTrees(i+1, e)
		sum += right * left
	}

	return sum

}
