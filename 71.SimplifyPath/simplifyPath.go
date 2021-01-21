package SimplifyPath

import "strings"

func simplifyPath(path string) string {
	stack := []string{""}
	buffer := []rune{}
	for _, c := range path {
		if c == '/' {
			stack = processBuffer(buffer, stack)
			buffer = buffer[:0]
			continue
		}
		buffer = append(buffer, c)
	}
	stack = processBuffer(buffer, stack)
	if len(stack) == 1 {
		return "/"
	}
	return strings.Join(stack, "/")
}

func processBuffer(buffer []rune, stack []string) []string {
	if len(buffer) == 0 || string(buffer) == "." {
		return stack
	}

	subPath := string(buffer)
	if subPath == ".." {
		if len(stack) == 1 {
			return stack
		}
		stack = stack[:len(stack)-1]
		return stack
	}
	stack = append(stack, subPath)
	return stack
}
