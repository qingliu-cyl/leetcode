package main

func findRepeatNumber(nums []int) int {
	resMap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := resMap[num]; ok {
			return num
		} else {
			resMap[num] = struct{}{}
		}
	}
	return 0
}

func main() {

}
