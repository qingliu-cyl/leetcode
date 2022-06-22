package main

import (
	"fmt"
	"strconv"
)

func minNumber(nums []int) string {
	min := func(i, j int) int {
		a, b := nums[i], nums[j]
		if a == 0 {
			return i
		}

		if b == 0 {
			return j
		}

		ab := a
		base := b
		for b != 0 {
			ab *= 10
			b /= 10
		}
		ab += base

		ba := base
		base = a
		for a != 0 {
			ba *= 10
			a /= 10
		}
		ba += base

		if ab > ba {
			return j
		}
		return i
	}

	for i := 0; i < len(nums)-1; i++ {
		idx := i
		for j := i+1; j < len(nums); j++ {
			idx = min(idx, j)
		}
		num := nums[i]
		nums[i] = nums[idx]
		nums[idx] = num
	}

	res := ""
	for _, i := range nums {
		res += strconv.Itoa(i)
	}

	return res
}

func main()  {
	fmt.Println(minNumber([]int{0,1}))
}
