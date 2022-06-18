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

func isStraight(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	sort(nums)

	king := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			king++
		}
	}

	pre := -1
	if nums[0] != 0 {
		pre = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			if ! (pre != -1 && nums[i] - pre <= king+1 || pre == -1) || nums[i] == pre{
				return false
			}

			if pre != -1 {
				king = king - (nums[i] - pre -1)
			}
			pre = nums[i]
		}
	}

	return true
}

func main()  {
	fmt.Println(isStraight([]int{8,2,9,7,10}))
}
