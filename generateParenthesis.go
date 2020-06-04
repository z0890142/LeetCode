package main

func generateParenthesis(n int) []string {
	var result []string
	rn, ln := n, n
	createParenthesis(rn, ln, "", &result)
	return result
}

func createParenthesis(rn int, ln int, preText string, result *[]string) {
	if ln == 0 && rn == 0 {
		*result = append(*result, preText)
	}
	if ln < rn || ln < 0 || rn < 0 {
		return
	}

	createParenthesis(rn-1, ln, preText+"(", result)
	createParenthesis(rn, ln-1, preText+")", result)

}
