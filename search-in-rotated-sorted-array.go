package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

func main() {

}
