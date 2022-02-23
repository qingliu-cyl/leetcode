package main

import "fmt"

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func nextPermutation(nums []int) {
	// 从右往左找最小数 nums[i] < num[i+1]
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 2
	isUp := true // 从右往左是否是一直增加， 如果是则说明已经是最后一个， 此时重新排序即可
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			isUp = false
			break
		}
	}

	if isUp {
		// 排序即可
		reverse(nums)
		return
	}

	// 查找从右往左第一个大于nums[i]的数
	j := len(nums) - 1
	for ; j > i; j-- {
		if nums[i] < nums[j] {
			break
		}
	}

	// 交换nums[i] nums[j]
	num := nums[i]
	nums[i] = nums[j]
	nums[j] = num

	// 将nums[i]之后的数排序
	reverse(nums[i+1:])

}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
