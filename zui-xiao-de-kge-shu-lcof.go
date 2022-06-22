package main

import "github.com/qingliu-cyl/golibs"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}

	min := func(i, j int) int {
		if arr[i] < arr[j] {
			return i
		}

		return j
	}

	for i := 0; i < k; i++{
		idx := i
		for j := i+1; j < len(arr); j++ {
			idx = min(idx, j)
		}
		num := arr[i]
		arr[i] = arr[idx]
		arr[idx] = num
	}

	return arr[:k]
}