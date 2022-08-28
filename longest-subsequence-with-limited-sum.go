package main

import (
	"fmt"
)

func sort(nums []int) {
	merge := func(l, m, r int) {
		list := make([]int, len(nums))
		copy(list, nums[:r+1])
		s, i, j := l, l, m+1
		for i <= m || j <= r {
			if i <= m && j <= r && list[i] < list[j] {
				nums[s] = list[i]
				i++
			} else if i <= m && j <= r && list[i] >= list[j] {
				nums[s] = list[j]
				j++
			} else if i > m && j <= r {
				nums[s] = list[j]
				j++
			} else if i <= m && j > r {
				nums[s] = list[i]
				i++
			} else {
				return
			}
			s++
		}
	}

	var mergeSort func(l, r int)
	mergeSort = func(l, r int) {
		if l < r {
			m := (l + r) / 2
			mergeSort(l, m)
			mergeSort(m+1, r)
			merge(l, m, r)
		}
	}
	mergeSort(0, len(nums)-1)
}

func maxLen(nums []int, query int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if query < nums[i] {
			break
		}
		query -= nums[i]
		res++
	}

	return res
}

func answerQueries(nums []int, queries []int) []int {
	answer := make([]int, len(queries))
	sort(nums)
	for i := 0; i < len(queries); i++ {
		answer[i] = maxLen(nums, queries[i])
	}

	return answer
}

func main() {
	fmt.Println(answerQueries([]int{4,5,2,1}, []int{3,10,21}))
}
