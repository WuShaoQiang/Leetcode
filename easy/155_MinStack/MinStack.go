package MinStack

import (
	"math"
)

/*
 Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

    push(x) -- Push element x onto stack.
    pop() -- Removes the element on top of the stack.
    top() -- Get the top element.
    getMin() -- Retrieve the minimum element in the stack.

Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

// Runtime: 60 ms, faster than 11.14% of Go online submissions for Min Stack.
// Memory Usage: 8.2 MB, less than 48.16% of Go online submissions for Min Stack.

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
