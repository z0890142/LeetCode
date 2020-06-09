package main

import "fmt"

func rotate(matrix [][]int) {

	for row := 0; row < (len(matrix) - 1); row++ {
		maxCol := len(matrix) - row - 1
		for col := 0; col < maxCol; col++ {
			tmp := matrix[row][col]
			matrix[row][col] = matrix[len(matrix)-1-col][maxCol]
			matrix[len(matrix)-1-col][maxCol] = tmp
		}
	}
	//上下交換
	for row := 0; row < len(matrix)/2; row++ {
		tmpRow := matrix[row]
		matrix[row] = matrix[len(matrix)-1-row]
		matrix[len(matrix)-1-row] = tmpRow
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
