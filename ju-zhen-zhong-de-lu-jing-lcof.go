package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	if len(board[0]) == 0 {
		return false
	}
	cow := len(board)
	cel := len(board[0])

	deepCopy := func(src [][]byte) [][]byte {
		res := make([][]byte, cow)
		for i := 0; i < cow; i++ {
			dst := make([]byte, cel)
			copy(dst, src[i])
			res[i] = dst
		}

		return res
	}


	var find  func(newBoard [][]byte, i, j int, idx int) bool
	find = func(newBoard [][]byte, i, j int, idx int) bool {
		if i >= cow || j >= cel || i < 0 || j < 0 || newBoard[i][j] != word[idx]{
			return false
		}

		if idx == len(word) - 1 {
			return true
		}

		copyBoard := deepCopy(newBoard)
		copyBoard[i][j] = 0
		return find(copyBoard, i+1, j, idx+1) || find(copyBoard, i-1, j, idx+1) ||
			find(copyBoard, i , j+1, idx+1) || find(copyBoard, i, j-1, idx+1)
	}

	for i := 0; i < cow; i++ {
		for j := 0; j < cel; j++ {
			if find(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{[]byte("a"), []byte("b"), []byte("ADEE")}, "ABCCED"))
}
