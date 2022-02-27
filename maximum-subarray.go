package main

// 最大子数组和
// https://leetcode-cn.com/problems/maximum-subarray/
// 动态规划法
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
