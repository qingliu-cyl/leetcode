package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	// 遍历， 至少有一个数会小于0
	for i := 0; i <= len(nums)-1; i++ {
		// 与上一个相同就不需要比较了
		if (i-1) >= 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(res)-float64(target)) {
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}

	}
	return res
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
