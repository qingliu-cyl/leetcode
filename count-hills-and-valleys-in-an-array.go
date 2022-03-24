package main

import "fmt"

func countHillValley(nums []int) int {
	// 去重
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	res := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] ||
			nums[i-1] > nums[i] && nums[i+1] > nums[i] {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5}))
}
