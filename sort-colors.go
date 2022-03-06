package main

import "fmt"

func sortColors(nums []int) {
	// 发现red就放到最前面， 发现blue就放到最后面， 发现white就不动
	foremost := 0
	last := len(nums) - 1
	for i := 0; i <= last; i++ {
		if nums[i] == 0 {
			nums[i] = nums[foremost]
			nums[foremost] = 0
			foremost++
		} else if nums[i] == 2 {
			nums[i] = nums[last]
			nums[last] = 2
			last--

			// 不等于1还需要再次判断
			if nums[i] != 1 {
				i--
			}
		}
	}
}

func main() {
	num := []int{2, 0, 2, 1, 1, 0}
	sortColors(num)
	fmt.Println(num)
}
