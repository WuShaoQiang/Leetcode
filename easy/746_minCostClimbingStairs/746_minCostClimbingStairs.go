package minCostClimbingStairs

/*
 On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:

Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

Example 2:

Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

Note:

    cost will have a length in the range [2, 1000].
    Every cost[i] will be an integer in the range [0, 999].

*/

// Runtime: 4 ms, faster than 99.43% of Go online submissions for Min Cost Climbing Stairs.
// Memory Usage: 2.9 MB, less than 92.92% of Go online submissions for Min Cost Climbing Stairs.

func minCostClimbingStairs(cost []int) int {
	minCost0, minCost1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		minCost0, minCost1 = minCost1, min(minCost0, minCost1)+cost[i]
	}
	return min(minCost0, minCost1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}