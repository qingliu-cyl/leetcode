package main

import (
	"fmt"
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	start := int64(1)
	res := int64(0)

	sum := func(i, j int64) int64 {
		if (j-i)%2 == 0 {
			return (i + (j - 1)) * ((j - i) / 2)
		} else {
			return (i+(j-1))*((j-i)/2) + (i+j-1)/2
		}

	}
	for _, i := range nums {
		if start == int64(i) {
			start++
			continue
		}

		if start > int64(i) {
			continue
		}
		if int64(i)-start < int64(k) {
			res += sum(start, int64(i))
			k = k - (i - int(start))
			start = int64(i) + int64(1)

		} else {
			res += sum(start, start+int64(k))
			return res
		}
	}

	last := int64(nums[len(nums)-1])
	res += sum(last+1, last+int64(k)+1)
	return res
}

func main() {
	//fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	fmt.Println(minimalKSum([]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}, 35))
}
