package blncdparantheses

func balancer(s string) bool {
	var stack []string
	for _, x := range s {
		switch x {
		case '(', '{', '[':
			stack = append(stack, string(x))

			continue
		case ')', '}', ']':
			if stack[len(stack)-1] != "(" {
				if stack[len(stack)-1] != "{" {
					if stack[len(stack)-1] != "[" {
						return false
					}
				}
			}
			stack = pop(stack)
		}
	}
	return true
}

func pop(stack []string) []string {
	if len(stack)-1 > 0 {
		stack = stack[:len(stack)-1]
	} else {
		stack[0] = " "
	}
	return stack
}
