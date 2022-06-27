package main

import "fmt"

func singleNumbers(nums []int) []int {
	res, left, right, m := 0, 0, 0, 1 // 与 a^0 = a

	for _, num := range nums {
		res ^= num
	}

	for (m & res) == 0 {
		m <<= 1
	}

	for _, num := range nums {
		if (num & m) != 0 {
			left ^= num
		} else {
			right ^= num
		}
	}

	return []int{left, right}
}

func main() {
	fmt.Println(singleNumbers([]int{1,2,5,2}))
}