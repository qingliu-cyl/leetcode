package main

import "fmt"

func insterionSert(list []int) {
	if len(list) <= 1 {
		return
	}

	for i := 1; i < len(list); i++ {
		v := list[i]

		j := i - 1
		for ; j >= 0 && list[j] > v; j-- {
			list[j+1] = list[j]
		}

		list[j+1] = v
	}
}

func main() {
	nums := []int{3, 1, 4, 5, 6, 7, 4, 8}
	insterionSert(nums)
	fmt.Println(nums)
}
