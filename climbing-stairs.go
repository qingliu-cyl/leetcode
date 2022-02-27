package main

func climbStairs(n int) int {
	res := 0
	var climb func(step, sum int)
	climb = func(step, sum int) {
		if sum + step == n {
			res += 1
			return
		} else if sum + step > n {
			return
		}

		climb(1, sum + step)
		climb(2, sum + step)
	}

	climb(1, 0)
	climb(2,0)

	return res
}

