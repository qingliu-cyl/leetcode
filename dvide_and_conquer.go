package main

import "fmt"

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

func main() {
	nums := []int{1, 2, 5, 3, 7, 4, 9, 2, 6, 0, 8, 10, 13, 11, 16, 14, 18, 20}
	sort(nums)
	fmt.Println(nums)
}
