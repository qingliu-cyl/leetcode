package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	res := make([]int, len(prices))
	res[0] = 0
	minPrice := prices[0]

	for i:=1; i<len(prices);i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		res[i] = max(res[i-1], prices[i] - minPrice)
	}

	return res[len(res)-1]
}

func main()  {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
