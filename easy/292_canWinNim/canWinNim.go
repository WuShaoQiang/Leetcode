package canWinNim

/*
You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.

Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.

Example:

Input: 4
Output: false 
Explanation: If there are 4 stones in the heap, then you will never win the game;
             No matter 1, 2, or 3 stones you remove, the last stone will always be 
             removed by your friend.*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Nim Game.
// Memory Usage: 1.9 MB, less than 71.64% of Go online submissions for Nim Game.

// 只要我把对手至于4的倍数上，无论它怎么拿，我都可以把最后一个石头拿走
// 同理，如果一开始的石头个数是4的倍数，无论我怎么拿，最后一个都不是我拿
func canWinNim(n int) bool {
	return !(n%4 == 0)
}
