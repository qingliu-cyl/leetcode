package main

import "fmt"

func removeStars(s string) string {
	stack := make([]byte, len(s))
	p := 0

	flag := []byte("*")[0]
	for i := 0; i < len(s); i++ {
		if s[i] == flag {
			p--
		} else {
			stack[p] = s[i]
			p++
		}
	}

	return string(stack[:p])
}

func main() {
	fmt.Println(removeStars("erase*****"))
}
