package main

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max( dp[j] * (i-j), (i-j)*j))
		}
	}

	return dp[n]
}