package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	l := 0
	r := len(nums) - 1
	mid := (l + r) / 2
	for l < r && mid != l && mid != r {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}
	if target < nums[l] || nums[l] == target {
		return l
	} else if nums[l] < target && target < nums[r] || nums[r] == target {
		return r
	} else {
		return r + 1
	}
}

func main() {

}
