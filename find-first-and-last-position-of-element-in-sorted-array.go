package main

import "fmt"

func searchRange(nums []int, target int) []int {
	// 查找第一个大于等于target的位置
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	// 若要实现O(logn), 这里也应该二分
	for r = l + 1; r < len(nums) && nums[r] <= target; r++ {
	}

	return []int{l, r - 1}
}

func main() {

	fmt.Println(searchRange([]int{1}, 1))
}
