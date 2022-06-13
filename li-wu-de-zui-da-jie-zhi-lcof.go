package main

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	res := make([][]int, len(grid))
	cel := len(grid[0])
	for i := 0; i < len(grid); i++ {
		res[i] = make([]int, cel)
	}
	res[0][0] = grid[0][0]

	max := func(i , j int)  int{
		if i > j {
			return i
		} else {
			return j
		}
	}

	maxNum := res[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < cel; j++ {
			if i == 0 && j == 0 {
				continue
			}else if i == 0 {
				res[i][j] = res[i][j-1] + grid[i][j]
			} else if j == 0 {
				res[i][j] = res[i-1][j] + grid[i][j]
			} else {
				res[i][j] = max(res[i][j-1]+grid[i][j],
					res[i-1][j]+grid[i][j])
			}

			if res[i][j] > maxNum {
				maxNum = res[i][j]
			}
		}
	}

	return maxNum
}