package main

func canJump(nums []int) bool {
	furthest := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for idx, num := range nums {
		if furthest >= idx {
			furthest = max(furthest, idx+num)
		} else {
			return false
		}
	}

	return true
}
