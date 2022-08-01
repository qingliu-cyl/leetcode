package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	digit, start, count := 1,1,9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count  = 9 * start * digit
	}

	num := start + (n-1) / digit
	res, _ := strconv.Atoi(string(strconv.Itoa(num)[(n-1)%digit]))
	return res
}

func main() {
	fmt.Println(findNthDigit(11))
}
