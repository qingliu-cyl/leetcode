package main

// 递归法
func climbStairs1(n int) int {
	res := 0
	var climb func(step, sum int)
	climb = func(step, sum int) {
		if step+sum == n {
			if sum+step == n {
				if sum+step == n {
					res += 1
					return
				} else if sum+step > n {
					return
				}

				climb(1, sum+step)
				climb(2, sum+step)
			}

			climb(1, 0)
			climb(2, 0)
		}
	}
	return res
}

// 动态规划法
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	//res[i] 表示爬到第i阶的方式
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	res[2] = 2

	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
