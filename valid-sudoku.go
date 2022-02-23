package main

func isValidSudoku(board [][]byte) bool {
	// 记录每一行每个数字是否已经出现过
	rowNum := [9][9]bool{}
	// 记录每一列每个数字是否已经出现过
	colNum := [9][9]bool{}
	// 记录每一个九宫格每个数字是否已经出现过
	cellNum := [3][3][9]bool{}

	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			// 记录每一行的数字
			if rowNum[i][v-'1'] {
				return false
			} else {
				rowNum[i][v-'1'] = true
			}

			// 记录每一列的数字
			if colNum[j][v-'1'] {
				return false
			} else {
				colNum[j][v-'1'] = true
			}

			// 记录每一个九宫个中的数字
			if cellNum[i/3][j/3][v-'1'] {
				return false
			} else {
				cellNum[i/3][j/3][v-'1'] = true
			}
		}
	}
	return true
}

func main() {

}
