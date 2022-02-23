package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	min := len(strs[0])
	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}

	equal := func(num int) bool {
		for i := 0; i < (len(strs) - 1); i++ {
			if strs[i][num] != strs[i+1][num] {
				return false
			}
		}
		return true
	}

	j := 0
	for ; j < min; j++ {
		if !equal(j) {
			break
		}
	}

	return strs[0][:j]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
