package main

func combine(n int, k int) [][]int {
	var result [][]int
	for val := 1; val <= n; val++ {
		findCombine([]int{val}, val, &result, n, k)
	}
	return result
}

func findCombine(preCombin []int, startIndex int, result *[][]int, n int, k int) {
	if len(preCombin) == k {
		tmp := make([]int, len(preCombin))
		copy(tmp, preCombin)
		*result = append(*result, tmp)
		return
	}
	for val := startIndex + 1; val <= n; val++ {
		preCombin = append(preCombin, val)
		findCombine(preCombin, val, result, n, k)
		preCombin = preCombin[0 : len(preCombin)-1]
	}
}
