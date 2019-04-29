package MinStack

import (
	"math"
)

type MinStack struct {
	stack []int
	head  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		head:  -1,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack)-1 > this.head {
		this.head++
		this.stack[this.head] = x
		return
	}
	this.stack = append(this.stack, x)
	this.head++
}

func (this *MinStack) Pop() {
	this.head--
}

func (this *MinStack) Top() int {
	return this.stack[this.head]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	for i := 0; i <= this.head; i++ {
		if this.stack[i] < min {
			min = this.stack[i]
		}
	}
	return min
}
