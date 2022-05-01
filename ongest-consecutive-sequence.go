package main

import "fmt"

func longestConsecutive(nums []int) int {
	resMap := make(map[int]struct{})

	for _, i := range nums {
		resMap[i] = struct{}{}
	}

	max := 0
	lenMap := make(map[int]int)
	for i, _ := range resMap {
		l := 1
		j := i - 1
		for {
			if _, ok := lenMap[j]; ok {
				l += lenMap[j]
				goto BREAK
			}
			if _, ok := resMap[j]; ok {
				l++
				j--
				continue
			}
			BREAK:
			if l > max {
				max = l
			}
			lenMap[i] = l
			break
		}

	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}
