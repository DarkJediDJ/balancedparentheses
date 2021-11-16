package blncdparantheses

func balancer(s string) bool {
	var stack []string
	for _, x := range s {
		if x == '(' || x == '[' || x == '{' {
			stack = append(stack, string(x))
			continue
		}
		if len(stack) == 0 {
			return false
		}
		current_char := stack[len(stack)-1]
		switch x {
		case ')':
			if current_char != "(" {
				return false
			}
			if len(stack)-1 > 0 {
				stack = stack[:len(stack)-1]
			} else {
				stack[0] = " "
			}
		case '}':
			if current_char != "{" {
				return false
			}
			if len(stack)-1 > 0 {
				stack = stack[:len(stack)-1]
			} else {
				stack[0] = " "
			}
		case ']':
			if current_char != "[" {
				return false
			}
			if len(stack)-1 > 0 {
				stack = stack[:len(stack)-1]
			} else {
				stack[0] = " "
			}
		}
	}
	return true
}
