package easy

func IsValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if c == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if c == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if c == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
