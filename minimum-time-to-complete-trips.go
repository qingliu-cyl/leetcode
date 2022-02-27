package main

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	// 找出耗时最短的一辆车
	var minTime int64 = int64(time[0])
	for _, i := range time {
		if minTime > int64(i) {
			minTime = int64(i)
		}
	}

	// 最长耗时为minTime * totalTrips
	maxTime  := minTime * int64(totalTrips)

	var find2 func(t int64) int64
	find2 = func(t int64) int64 {
		neetReturn := false
		for _, i := range time {
			if t % int64(i) == 0 {
				neetReturn = true
			}
		}

		if neetReturn {
			return t
		} else {
			return find2(t-1)
		}
	}

	var find func(l, r int64) int64
	find = func(l, r int64) int64 {
		if l == r {
			return find2(l)
		}
		mid := (r + l) / 2
		var total int64 = 0
		for _, i := range time {
			total += mid / int64(i)
		}

		if total == int64(totalTrips) {
			return find2(mid)
		}

		if total < int64(totalTrips) {
			l = mid+1
		} else {
			r = mid
		}
		return find(l, r)
	}

	return find(0, maxTime)
}

func main() {
	fmt.Println(minimumTime([]int{3,3,8}, 6))
}
