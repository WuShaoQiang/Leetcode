package maxDistToClosest

/*
In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

Example 1:

Input: [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.

Example 2:

Input: [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat, the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.

Note:

    1 <= seats.length <= 20000
    seats contains only 0s or 1s, at least one 0, and at least one 1.


*/

// Runtime: 12 ms, faster than 100.00% of Go online submissions for Maximize Distance to Closest Person.
// Memory Usage: 5.4 MB, less than 66.67% of Go online submissions for Maximize Distance to Closest Person.

func maxDistToClosest(seats []int) int {
	start := 0
	max := 0
	var i int
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			tmp := i - start
			if start != 0 {
				tmp = (tmp + 1) / 2
			}
			if max < tmp {
				max = tmp
			}
			start = i + 1
		}
	}
	if max < i-start {
		max = i - start
	}

	return max
}
