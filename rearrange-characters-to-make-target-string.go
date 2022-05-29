package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	targetMap := make(map[int32]int)
	for _, v := range target {
		targetMap[v] += 1
	}

	for _, v := range s {
		if _, ok := targetMap[v]; ok {
			targetMap[v]++
		}
	}

	for _, v := range target {
		targetMap[v]--
	}
	res := 0
	for {
		for _, v := range target {
			if num, ok := targetMap[v];ok {
				if num != 0 {
					targetMap[v]--
				} else {
					goto OUT
				}
			} else {
				goto OUT
			}
		}
		res++
	}
	OUT:
		return res
}

func main() {
	fmt.Println(rearrangeCharacters("abbaccaddaeea", "aaaaa"))
}
