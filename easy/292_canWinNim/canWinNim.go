package canWinNim

// 只要我把对手至于4的倍数上，无论它怎么拿，我都可以把最后一个石头拿走
// 同理，如果一开始的石头个数是4的倍数，无论我怎么拿，最后一个都不是我拿
func canWinNim(n int) bool {
	return !(n%4 == 0)
}
