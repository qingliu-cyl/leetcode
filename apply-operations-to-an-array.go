package main

import "fmt"

func applyOperations(nums []int) []int {
	res := make([]int, len(nums))
	p := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 && nums[i] == nums[i+1] {
			res[p] = nums[i] * 2
			p++
			nums[i+1] = 0
		} else if nums[i] != 0{
			res[p] = nums[i]
			p++
		}
	}

	if nums[len(nums)-1] != 0 {
		res[p] = nums[len(nums)-1]
	}

	return res
}

func main() {
	fmt.Println(applyOperations([]int{0,1}))
}
