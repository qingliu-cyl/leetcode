package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	last := nums[0]
	for i := 1; i < len(nums); {
		if nums[i] == last {
			if i != len(nums)-1 {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
		} else {
			last = nums[i]
			i++
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
