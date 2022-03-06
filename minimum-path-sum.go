package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0]) // 题目中m,n 均大于1， 所以不需考虑数组越界
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = grid[0][i] + dp[0][i-1]
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
