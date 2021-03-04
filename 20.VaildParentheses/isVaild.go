package VaildParentheses

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		switch char := s[i]; char {
		case '(', '[', '{':
			stack = append(stack, char)
		default:
			if len(stack) == 0 {
				return false
			}

			peek := stack[len(stack)-1]

			if peek == '(' && char != ')' {
				return false
			}

			if peek == '[' && char != ']' {
				return false
			}

			if peek == '{' && char != '}' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
