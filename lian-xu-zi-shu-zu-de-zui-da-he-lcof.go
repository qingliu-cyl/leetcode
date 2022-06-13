package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := make([]int, len(nums))

	max := func(i , j int)  int{
		if i > j {
			return i
		} else {
			return j
		}
	}
	res[0] = nums[0]
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = max(res[i-1]+nums[i], nums[i])
		if res[i] > maxNum {
			maxNum = res[i]
		}
	}

	return  maxNum
}
