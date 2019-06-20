package isRobotBounded

/*
On an infinite plane, a robot initially stands at (0, 0) and faces north.  The robot can receive one of three instructions:

    "G": go straight 1 unit;
    "L": turn 90 degrees to the left;
    "R": turn 90 degress to the right.

The robot performs the instructions given in order, and repeats them forever.

Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.



Example 1:

Input: "GGLLGG"
Output: true
Explanation:
The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.

Example 2:

Input: "GG"
Output: false
Explanation:
The robot moves north indefinitely.

Example 3:

Input: "GL"
Output: true
Explanation:
The robot moves from (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...



Note:

    1 <= instructions.length <= 100
    instructions[i] is in {'G', 'L', 'R'}


*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Robot Bounded In Circle.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Robot Bounded In Circle.

func isRobotBounded(instructions string) bool {
	direction := 0
	res := []int{0, 0}
	for cnt := 0; cnt < 4; cnt++ {
		for i := 0; i < len(instructions); i++ {
			if instructions[i] == 'G' {
				switch {
				case direction%4 == 0:
					res[1]++
				case direction%4 == 1 || direction%4 == -3:
					res[0]++
				case direction%4 == 2 || direction%4 == -2:
					res[1]--
				case direction%4 == 3 || direction%4 == -1:
					res[0]--
				}
			} else if instructions[i] == 'L' {
				direction--
			} else {
				direction++
			}
		}
		if res[0] == 0 && res[1] == 0 {
			return true
		}
	}

	return res[0] == 0 && res[1] == 0
}
