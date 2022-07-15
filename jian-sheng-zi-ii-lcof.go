package main

// 贪心
func cuttingRope(n int) int {
	if n <= 3 {
		return n-1
	}

	ret := 1
	if n % 3 == 2 {
		ret = 2
		n -= 2
	}

	if n % 3 == 1 {
		ret = 4
		n -= 4
	}

	for n > 0 {
		ret = ret * 3 % 1000000007
		n -= 3
	}

	return ret
}