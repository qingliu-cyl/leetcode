package main

import (
	"fmt"
)

//从栈中取出 K 个硬币的最大面值和
// 超时了
func maxValueOfCoins(piles [][]int, k int) int {
	res := 0
	l := len(piles)

	var search func(addr []int, num int, times int)
	search = func(addr []int, num int, times int) {
		if times == 0 {
			if num > res {
				res = num
			}
			return
		}
		for idx, _ := range addr {
			if addr[idx] < len(piles[idx]) {
				newAddr := make([]int, l)
				copy(newAddr, addr)
				newAddr[idx]++
				search(newAddr, num+piles[idx][addr[idx]], times-1)
			}
		}
	}

	addr := make([]int, l)
	search(addr, 0, k)
	return res
}

func main() {
	fmt.Println(maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))
}
