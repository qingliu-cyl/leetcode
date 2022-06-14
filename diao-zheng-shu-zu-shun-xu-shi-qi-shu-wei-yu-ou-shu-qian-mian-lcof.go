package main

func exchange(nums []int) []int {
	for j := 0; j < len(nums); j++ {
		if nums[j] % 2 != 0 {
			num := nums[j]
			nums = append(nums[:j], nums[j+1:]...)
			nums = append([]int{num}, nums...)
		}
	}

	return nums
}