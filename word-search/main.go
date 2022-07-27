package main

var (
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

func backtrack(i, j int, board [][]byte, word string, matched int) bool {
	if matched == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[matched] {
		return false
	}
	t := board[i][j]
	board[i][j] = '#'
	for k := 0; k < 4; k++ {
		if backtrack(i+dx[k], j+dy[k], board, word, matched+1) {
			return true
		}
	}
	board[i][j] = t
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(i, j, board, word, 0) {
				return true
			}
		}
	}
	return false
}
