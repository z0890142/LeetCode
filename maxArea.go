package main

func maxArea(height []int) int {
	leftIndex := 0
	rightIndex := len(height) - 1

	var heightVal int
	var MaxArea int

	for leftIndex < rightIndex {
		if height[leftIndex] < height[rightIndex] {
			heightVal = height[leftIndex]
		} else {
			heightVal = height[rightIndex]
		}

		tmpArea := heightVal * (rightIndex - leftIndex)
		if tmpArea > MaxArea {
			MaxArea = tmpArea
		}
		if height[leftIndex] < height[rightIndex] {
			leftIndex += 1
		} else {
			rightIndex -= 1
		}
	}
	return MaxArea
}
