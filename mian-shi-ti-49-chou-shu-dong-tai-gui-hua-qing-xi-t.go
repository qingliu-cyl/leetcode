package main

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	a,b,c := 0,0,0

	min := func(a, b, c int) int {
		if a >= c && b >= c {
			return c
		}
		if a >= b && c >= b {
			return b
		}
		return a
	}

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a] * 2, dp[b]*3, dp[c]*5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}

	return dp[n-1]
}

func main()  {
	nthUglyNumber(10)
}