package NumberofIslands

func numIslands(grid [][]byte) int {
	var result int
	for rIndex, row := range grid {
		for cIndex, col := range row {
			if col == 'x' || col == '0' {
				continue
			}

			if col == '1' {
				result++
				numIslandsHelper(&grid, rIndex, cIndex)
			}
		}
	}
	return result
}

func numIslandsHelper(grid *[][]byte, rIndex int, cIndex int) {
	(*grid)[rIndex][cIndex] = 'x'

	if rIndex-1 >= 0 && (*grid)[rIndex-1][cIndex] == '1' {
		numIslandsHelper(grid, rIndex-1, cIndex)
	}

	if rIndex+1 < len(*grid) && (*grid)[rIndex+1][cIndex] == '1' {
		numIslandsHelper(grid, rIndex+1, cIndex)
	}

	if cIndex-1 >= 0 && (*grid)[rIndex][cIndex-1] == '1' {
		numIslandsHelper(grid, rIndex, cIndex-1)
	}

	if cIndex+1 < len((*grid)[0]) && (*grid)[rIndex][cIndex+1] == '1' {
		numIslandsHelper(grid, rIndex, cIndex+1)
	}

}
