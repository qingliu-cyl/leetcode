package main

func search(nums []int, target int) int {
	res := 0

	for _, num := range nums {
		if num == target {
			res++
		}
	}

	return res
}

func main() {

}
