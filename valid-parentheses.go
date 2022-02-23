package main

import "fmt"

var dict = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isValid(s string) bool {
	stack := make([]string, len(s))
	ptr := 0 // 栈指针

	for _, i := range s {
		if ptr == 0 || stack[ptr-1] != dict[string(i)] {
			stack[ptr] = string(i)
			ptr++
		} else if stack[ptr-1] == dict[string(i)] {
			ptr--
		}
	}

	return ptr == 0
}

func main() {
	s := "aaasdsds"
	for _, i := range s {
		fmt.Println(string(i))
	}
}
