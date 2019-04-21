package isValid

import "container/list"

type Stack struct {
	List *list.List
}

func isValid(s string) bool {
	stack := Stack{List: list.New()}
	for _, v := range s {
		elem := stack.List.Back()
		if elem == nil {
			stack.List.PushBack(string(v))
			continue
		}
		switch elem.Value.(string) {
		case `(`:
			if string(v) != `)` {
				goto Push
			}
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
