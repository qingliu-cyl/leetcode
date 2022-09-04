package main

import "fmt"

func numberOfWays(startPos int, endPos int, k int) int {
	res := 0
	NUM := 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10  + 7

	var recursion func(point, steps int)
	recursion = func(point, steps int) {
		if endPos - point > steps {
			return
		}
		if steps == 0 {
			if point == endPos {
				res = (res + 1) %  NUM
			}
			return
		}

		recursion(point-1, steps-1)
		recursion(point+1, steps-1)
	}

	recursion(startPos-1, k-1)
	recursion(startPos+1, k-1)

	return res
}

func main() {
	fmt.Println(numberOfWays(264, 198, 68))
}
