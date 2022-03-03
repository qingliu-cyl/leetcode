package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	max := func(i, j int) int {
		if i < j {
			return j
		} else {
			return i
		}
	}

	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1] = []int{res[len(res)-1][0],
				max(res[len(res)-1][1], intervals[i][1])}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
