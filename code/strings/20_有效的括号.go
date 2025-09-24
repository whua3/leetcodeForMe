package strings

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
