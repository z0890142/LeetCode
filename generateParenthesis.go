package main

func generateParenthesis(n int) []string {
	var result []string
	ln, rn := n, n
	Parenthesis(ln, rn, "", &result)
	return result
}

func Parenthesis(left int, right int, text string, result *[]string) {
	if right < 0 || left < 0 || left > right {
		return
	}
	if right == 0 && left == 0 {
		*result = append(*result, text)
	}
	Parenthesis(left, right-1, text+")", result)
	Parenthesis(left-1, right, text+"(", result)

}
