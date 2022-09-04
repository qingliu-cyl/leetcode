package main

import "fmt"

func longestNiceSubarray(nums []int) int {
	max := 1

	var nice func(child []int, num int) bool
	nice = func(child []int, num int) bool {
		for i := 0; i < len(child); i++ {
			if child[i]&num != 0 {
				return false
			}
		}

		return true
	}
	s, end := 0, 0
	for i := 1; i < len(nums); {
		if nice(nums[s:end+1], nums[i]) {
			end++
			i++
		} else {
			if end-s+1 > max {
				max = end - s + 1
			}
			s++
			if s == i {
				end = s
				i++
			}
		}
	}

	if end-s+1 > max {
		max = end - s + 1
	}

	return max

}

func main() {
	fmt.Println(longestNiceSubarray([]int{135745088,609245787,16,2048,2097152}))
}
