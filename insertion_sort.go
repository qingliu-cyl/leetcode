package main

import "fmt"

// 插入排序
func insterionSert(list []int) {
	if len(list) <= 1 {
		return
	}

	for i := 1; i < len(list); i++ {
		v := list[i]

		j := i - 1
		for ; j >= 0 && list[j] > v; j-- {
			list[j+1] = list[j]
		}

		list[j+1] = v
	}
}

// 插入排序2
func insertSort(nums []int){
	tmp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		j := i-1
		for j >= 0  && nums[j] > nums[i]{
			j--
		}

		copy(tmp, nums)
		nums[j+1] = nums[i]
		copy(nums[j+2:i+1], tmp[j+1:i])
	}
}

func main() {
	nums := []int{3, 1, 4, 5, 6, 7, 4, 8}
	insterionSert(nums)
	fmt.Println(nums)
}
