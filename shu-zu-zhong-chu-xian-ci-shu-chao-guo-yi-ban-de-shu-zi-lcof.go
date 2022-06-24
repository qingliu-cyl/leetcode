package main

func majorityElement(nums []int) int {
	num := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count = 1
			num = nums[i]
			continue
		}

		if nums[i] == num {
			count++
			continue
		}

		if count > 0 {
			count--
		}
	}

	return num
}

func main()  {
	majorityElement([]int{1,2,3,2,2,2,5,4,2})
}