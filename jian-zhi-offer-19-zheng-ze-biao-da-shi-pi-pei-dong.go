package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 2; i <= n; i++ {
		dp[0][i] = dp[0][i-2]  && p[i-1] == []byte("*")[0]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == []byte("*")[0] {
				if dp[i][j-2] {
					dp[i][j] = dp[i][j-2]
				} else if dp[i-1][j] &&
					( s[i-1] == p[j-2] || p[j-2] == []byte(".")[0]) {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j-1] &&
					( s[i-1] == p [j-1] || p[j-1] == []byte(".")[0]) {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aaa", "a*"))
}
