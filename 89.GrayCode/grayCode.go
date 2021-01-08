package GrayCode

func grayCode(n int) []int {
	if n < 0 {
		return nil
	}

	res := []int{0}
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]|(1<<uint(i)))
		}
	}

	return res
}
