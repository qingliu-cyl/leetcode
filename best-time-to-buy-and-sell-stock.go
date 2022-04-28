package main

func maxProfit(prices []int) int {
	minPrice := 10 * 10 * 10 * 10
	maxProfit := 0
	for _, i := range prices {
		if minPrice > i {
			minPrice = i
		} else {
			profit := i - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func main() {

}
