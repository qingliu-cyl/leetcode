package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); {
		if nums[i] == val {
			if i != len(nums)-1 {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
		} else {
			val = nums[i]
			i++
		}
	}
	return len(nums)
}

func main() {

}
