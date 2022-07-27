package main

import "fmt"

func pow(x, n int) int {
	ans := 1
	for n != 0 {
		ans *= 2
		n--
	}
	return ans * x
}

func singleNumber(nums []int) int {
	count := make([]int, 32)

	for idx, _ := range nums {
		for  i := 0; i < 32; i++ {
			count[i] +=  nums[idx] & 1
			nums[idx] = nums[idx]>>1
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		count[i] = count[i] % 3
		res += pow(count[i], i)
	}

	return res
}

func main() {
	fmt.Println(singleNumber([]int{3,4,3,3}))
}
