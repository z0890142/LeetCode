package WordSearch

func exist(board [][]byte, word string) bool {
	for rIndex, r := range board {
		for cIndex, _ := range r {
			if existHelper(board, rIndex, cIndex, word) {
				return true
			}
		}
	}

	return false
}

func existHelper(board [][]byte, rIndex int, cIndex int, word string) bool {
	if word == "" {
		return true
	}
	if rIndex < 0 || rIndex > len(board)-1 || cIndex < 0 || cIndex > len(board[0])-1 {
		return false
	}
	if board[rIndex][cIndex] != word[0] || board[rIndex][cIndex] == '#' {
		return false
	}
	tmp := board[rIndex][cIndex]
	board[rIndex][cIndex] = '#'

	if existHelper(board, rIndex-1, cIndex, word[1:]) {
		return true
	}
	if existHelper(board, rIndex+1, cIndex, word[1:]) {
		return true
	}
	if existHelper(board, rIndex, cIndex-1, word[1:]) {
		return true
	}
	if existHelper(board, rIndex, cIndex+1, word[1:]) {
		return true
	}
	board[rIndex][cIndex] = tmp
	return false
}
