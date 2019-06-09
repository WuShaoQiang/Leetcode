package divisorGame

/*
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:

    Choosing any x with 0 < x < N and N % x == 0.
    Replacing the number N on the chalkboard with N - x.

Also, if a player cannot make a move, they lose the game.

Return True if and only if Alice wins the game, assuming both players play optimally.



Example 1:

Input: 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.

Example 2:

Input: 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.



Note:

    1 <= N <= 1000


*/

func divisorGame(N int) bool {
	// 一开始是奇数的话，Alice肯定会将它变成偶数
	// 所以Bob会继续将偶数变为奇数，利用N%1=0这个规则，N=N-1就会变成奇数
	// 到最后Alice将3变为2，Bob就赢了

	//所以只用判断一开始的数是偶数还是奇数
	return N%2 == 0
}
