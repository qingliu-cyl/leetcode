package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	head, end := "", ""
	for i := 0; i < len(s); i++ {
		if i < n {
			head += string(s[i])
		} else {
			end += string(s[i])
		}
	}

	return end + head
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}
