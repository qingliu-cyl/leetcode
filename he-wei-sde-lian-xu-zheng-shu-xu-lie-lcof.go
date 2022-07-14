package main

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)

	l,r, sum :=1,2,3
	add := func() {
		list := make([]int, 0)
		for i := l; i <= r; i++ {
			list = append(list, i)
		}
		res = append(res, list)
	}

	for l < r {
		if sum == target {
			add()
		}
		if sum < target {
			r++
			sum += r
		} else {
			sum -= l
			l++
		}
	}

	return res
}