package main

import "fmt"

func maxArea(height []int) int {
	minFUnc := func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}

	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		size := (r - l) * minFUnc(height[l], height[r])
		if size > res {
			res = size
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
