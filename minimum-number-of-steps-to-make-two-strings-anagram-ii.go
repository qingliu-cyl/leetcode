package main

import "fmt"

func minSteps(s string, t string) int {
	mapS := make(map[string]int)
	for _, i := range s {
		if _, ok := mapS[string(i)];ok {
			mapS[string(i)]++
		} else {
			mapS[string(i)] = 1
		}
	}

	mapT := make(map[string]int)
	for _, i := range t {
		if _, ok := mapT[string(i)];ok {
			mapT[string(i)]++
		} else {
			mapT[string(i)] = 1
		}
	}
	tAddNum := 0
	for _, i := range s {
		if v, ok := mapT[string(i)];ok && v > 0 {
			mapT[string(i)]--
		} else {
			tAddNum++
		}
	}

	sAddSum := 0
	for _, i := range t {
		if v, ok := mapS[string(i)];ok && v > 0 {
			mapS[string(i)]--
		} else {
			sAddSum++
		}
	}
	return tAddNum + sAddSum
}

func main()  {
	fmt.Println(minSteps("leetcode", "coats"))
}
