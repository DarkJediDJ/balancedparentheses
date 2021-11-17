package blncdparantheses

func balancer(s string) bool {
	var stack []rune
	parentheses := [3]rune{'(', '[', '{'}
loop:
	for _, x := range s {
		switch x {
		case '(', '{', '[':
			stack = append(stack, x)

			continue
		case ')', '}', ']':
			for _, i := range parentheses {
				if stack[len(stack)-1] == i {
					stack = pop(stack)
					continue loop
				}
			}
			return false
		}
	}
	return true
}

func pop(stack []rune) []rune {
	if len(stack)-1 > 0 {
		stack = stack[:len(stack)-1]
	} else {
		stack[0] = ' '
	}
	return stack
}
