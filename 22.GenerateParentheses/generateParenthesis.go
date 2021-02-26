package GenerateParentheses

func generateParenthesis(n int) []string {
	var result []string
	rn, ln := n, n
	generateParenthesisHelper(ln, rn, "", &result)
	return result
}

func generateParenthesisHelper(l int, r int, preText string, result *[]string) {
	if r == 0 && l == 0 {
		*result = append(*result, preText)
	}
	if r < l || r < 0 || l < 0 {
		return
	}
	generateParenthesisHelper(l-1, r, preText+"(", result)
	generateParenthesisHelper(l, r-1, preText+")", result)

}
