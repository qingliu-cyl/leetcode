package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

func uniquePaths1(m int, n int) int {
	res := 0

	var dfs func(m, n int)
	dfs = func(m, n int) {
		if m == 0 || n == 0 {
			res++
			return
		}

		dfs(m-1, n)
		dfs(m, n-1)
	}
	dfs(m-1, n-1)
	return res
}

func main() {
	fmt.Println(uniquePaths(51, 9))
}
