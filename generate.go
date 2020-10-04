package main

func generate(numRows int) [][]int {
	var result [][]int
	result = append(result, []int{1}, []int{1, 1})
	if numRows < 3 {
		return result[:numRows]
	}
	for index := 1; index < numRows-1; index++ {
		var tmpSlice []int
		for index2 := 0; index2 < len(result[index])-1; index2++ {
			tmpSlice = append(tmpSlice, (result[index][index2] + result[index][index2+1]))
		}
		tmpSlice = append(tmpSlice, 1)
		tmpSlice = append([]int{1}, tmpSlice...)
		result = append(result, tmpSlice)
	}
	return result
}
