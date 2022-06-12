package main

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	res := make([]int, n+1)
	res[0], res[1], res[2] = 1,1,2
	for i := 3; i <=n; i++ {
		res[i] = (res[i-1] + res[i-2])%1000000007
	}
	return res[n]
}