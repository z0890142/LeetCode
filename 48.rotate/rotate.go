package rotate

import "fmt"

func rotate(matrix [][]int) {

	for row := 0; row < (len(matrix) - 1); row++ {
		maxCol := len(matrix[row]) - row - 1
		for col := 0; col < maxCol; col++ {
			matrix[row][col], matrix[len(matrix)-col-1][maxCol] = matrix[len(matrix)-col-1][maxCol], matrix[row][col]
		}
	}
	for row := 0; row < len(matrix)/2; row++ {
		matrix[row], matrix[len(matrix)-row-1] = matrix[len(matrix)-row-1], matrix[row]
	}
}

func show(matrix [][]int) {

	for index, _ := range matrix {
		for col := 0; col < len(matrix); col++ {
			fmt.Print(matrix[index][col], ",")
		}
		fmt.Print("\n")

	}
}
