package main

import "fmt"

func mostFrequentEven(nums []int) int {
	ret := make(map[int]int)
	for _, i := range nums {
		if _, ok := ret[i]; ok {
			ret[i]++
		}else {
			ret[i] = 1
		}
	}
	ret[-1] = 0

	max := -1
	for k, v := range ret {
		if k %2 == 0 && (v > ret[max] || v == ret[max] && k < max){
			max = k
		}
	}

	return max
}

func main() {
	fmt.Println(mostFrequentEven([]int{0,1,2,2,4,4,1}))
}