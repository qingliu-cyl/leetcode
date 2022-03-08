package main

import "fmt"

func exist(board [][]byte, word string) bool {
	res := false

	copyList := func(list [][]byte) [][]byte {
		newList := make([][]byte, len(list))
		for idx, _ := range newList {
			newList[idx] = make([]byte, len(list[idx]))
			copy(newList[idx], list[idx])
		}

		return newList
	}

	var dfs func(list [][]byte, x, y int, idx int)
	dfs = func(list [][]byte, x, y int, idx int) {
		if idx >= len(word) {
			res = true
			return
		}

		if x-1 >= 0 && list[x-1][y] == word[idx] {
			newList := copyList(list)
			newList[x-1][y] = 0
			dfs(newList, x-1, y, idx+1)
			if res {
				return
			}
		}
		if y-1 >= 0 && list[x][y-1] == word[idx] {
			newList := copyList(list)
			newList[x][y-1] = 0
			dfs(newList, x, y-1, idx+1)
			if res {
				return
			}
		}

		if x+1 < len(board) && list[x+1][y] == word[idx] {
			newList := copyList(list)
			newList[x+1][y] = 0
			dfs(newList, x+1, y, idx+1)
			if res {
				return
			}
		}

		if y+1 < len(board[0]) && list[x][y+1] == word[idx] {
			newList := copyList(list)
			newList[x][y+1] = 1
			dfs(newList, x, y+1, idx+1)
			if res {
				return
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			}

			newList := copyList(board)
			newList[i][j] = 0
			dfs(newList, i, j, 1)
			if res {
				return res
			}
		}
	}

	return res
}

func main() {
	l := [][]byte{[]byte("CAA"), []byte("AAA"), []byte("BCD")}
	fmt.Println(exist(l, "AAB"))
}
