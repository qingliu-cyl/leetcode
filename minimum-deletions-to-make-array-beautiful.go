package main

import "fmt"

// 美化数组的最少删除数
func minDeletion(nums []int) int {
	res := 0

	for i := 0; i < len(nums)-1; {
		if nums[i] != nums[i+1] {
			i += 2
		} else {
			i++
			res++
		}
	}

	if (len(nums)-res)%2 == 0 {
		return res
	} else {
		return res + 1
	}
}

func main() {
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
}
