package main

//Balancer checks string for balanced parentheses
func Balancer(s string) bool {
	var stack []rune
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, x := range s {
		switch x {
		case '(', '{', '[':
			stack = append(stack, x)

			continue
		case ')', '}', ']':
			if stack[len(stack)-1] != m[x] {
				return false
			}
			stack = pop(stack)
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
