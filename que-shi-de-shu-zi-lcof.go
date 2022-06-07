package main

func missingNumber(nums []int) int {
	target := 0
	for _, num := range nums {
		if num == target {
			target++
			continue
		} else {
			return target
		}
	}
	return target
}

func main() {

}
