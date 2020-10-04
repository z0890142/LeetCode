package main

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	tempPreNum := 0
	tempAfterNum := 0
	for i := 0; i <= rowIndex; i++ {
		if i >= 2 {
			tempPreNum = 1
			tempAfterNum = result[1]
		}
		for j := 0; j < i+1; j++ {
			if j == 0 {
				result[0] = 1
			} else if j == i {
				result[j] = 1
			} else {
				result[j] = tempPreNum + tempAfterNum
				tempPreNum = tempAfterNum
				tempAfterNum = result[j+1]
			}
		}
	}
	return result
}
