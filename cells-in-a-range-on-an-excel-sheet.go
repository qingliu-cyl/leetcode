package main

import "fmt"

func cellsInRange(s string) []string {
	celStart := rune(s[0])
	celEnd := rune(s[3])

	rowStart := rune(s[1])
	rowEnd := rune(s[4])

	res := make([]string, 0)

	for i := celStart; i <= celEnd; i++ {
		for j := rowStart; j <= rowEnd; j++ {
			res = append(res, string(i)+string(j))
		}
	}

	return res
}

func main() {
	fmt.Println(cellsInRange("K1:L2"))
}
