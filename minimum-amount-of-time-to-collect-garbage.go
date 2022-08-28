package main

import (
	"fmt"
)

func garbageCollection(garbage []string, travel []int) int {
	time := 0
	var target byte

	getTime := func() {
		p := 0
		for i := 0; i < len(garbage); i++ {
			for j := 0; j < len(garbage[i]); j++ {
				if target == garbage[i][j] {
					// 车不在当前位置：
					if p != i {
						for s := p; s < i; s++ {
							time += travel[s]
						}
						p = i
					}

					time++
				}

			}
		}
	}

	// 金属车
	target = []byte("M")[0]
	getTime()
	// 纸车
	target = []byte("P")[0]
	getTime()
	// 玻璃车
	target = []byte("G")[0]
	getTime()
	return time
}

func main() {
	fmt.Println(garbageCollection([]string{"MMM","PGM","GP"}, []int{3,10}))
}
