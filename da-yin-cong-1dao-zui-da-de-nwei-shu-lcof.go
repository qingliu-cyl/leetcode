package main

func printNumbers(n int) []int {
	max := 1
	for i := 0; i < n; i++ {
		max *=10
	}

	res := make([]int, max-1)
	for i := 1; i < max; i++ {
		res[i-1] = i
	}
	return res
}