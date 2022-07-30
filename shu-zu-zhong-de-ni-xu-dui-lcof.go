package main

import "fmt"

func reversePairs(nums []int) int {
	var mergeSort func(l, r int)
	res := 0
	tmp := make([]int, len(nums))
	mergeSort = func(l, r int) {
		if r <= l {
			return
		}

		m := (l+r)/ 2
		mergeSort(l, m)
		mergeSort(m+1, r)

		i, j := l, m+1
		copy(tmp[l:r+1], nums[l:r+1])
		for k := l; k <= r; k++ {
			if i == m+1 {
				//左边已经全部遍历结束， 直接赋值
				nums[k] = tmp[j]
				j++
			} else if j == r+1 || tmp[i] <= tmp[j]{
				// 右边遍历结束或者做小于右
				nums[k] = tmp[i]
				i++
			} else {
				// 左大于右场景
				nums[k] = tmp[j]
				j++
				res += (m-i+1)
			}
		}
	}
	mergeSort(0, len(nums)-1)
	return res
}

func main() {
	fmt.Println(reversePairs([]int{1,3,2,3,1}))
}
