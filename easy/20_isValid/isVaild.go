package isValid

import "container/list"

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true

Example 2:

Input: "()[]{}"
Output: true

Example 3:

Input: "(]"
Output: false

Example 4:

Input: "([)]"
Output: false

Example 5:

Input: "{[]}"
Output: true
*/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.2 MB, less than 27.46% of Go online submissions for Valid Parentheses.

type Stack struct {
	List *list.List
}

// 利用栈的特性,如果有匹配则直接pop出来,没有就push进去
func isValid(s string) bool {
	stack := Stack{List: list.New()}
	for _, v := range s {
		elem := stack.List.Back()
		// 如果栈里面没有元素,就直接往里面加
		if elem == nil {
			stack.List.PushBack(string(v))
			continue
		}
		// assertion
		switch elem.Value.(string) {
		case `(`:
			// 只要不符合,就push
			if string(v) != `)` {
				goto Push
			}
			// 符合就pop,这里的remove最后一个其实就是pop的作用
			stack.List.Remove(stack.List.Back())
		case `[`:
			if string(v) != `]` {
				goto Push
			}
			stack.List.Remove(stack.List.Back())
		case `{`:
			if string(v) != `}` {
				goto Push
			}
			stack.List.Remove(stack.List.Back())
		}
		continue
	Push:
		stack.List.PushBack(string(v))
	}
	return stack.List.Len() == 0
}
