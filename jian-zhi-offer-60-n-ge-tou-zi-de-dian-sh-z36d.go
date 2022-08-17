package main

import "fmt"

func dicesProbability(n int) []float64 {
	dp := []float64{float64(1)/6, float64(1)/6, float64(1)/6, float64(1)/6, float64(1)/6, float64(1)/6}

	for i := 2; i < n+1; i++ {
		tmp := make([]float64, 5 * i +1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / float64(6)
			}
		}

		dp = tmp
	}

	return dp
}

func main() {
	fmt.Println(dicesProbability(1))
}
