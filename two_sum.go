package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if targetIndex, ok := numMap[target-num]; ok {
			return []int{index, targetIndex}
		} else {
			numMap[num] = index
		}
	}
	return []int{}
}
