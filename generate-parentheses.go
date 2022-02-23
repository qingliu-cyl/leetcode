package main

import (
	"fmt"
)

var res []string
var dict = map[string]string{
	")": "(",
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

func backtrack(num, deep, l, r int, str string) {
	if deep == num*2-1 && isValid(str) {
		res = append(res, str)
	} else if l <= num && r <= num {
		backtrack(num, deep+1, l+1, r, str+"(")
		backtrack(num, deep+1, l, r+1, str+")")
	}
}

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	backtrack(n, 0, 1, 0, "(")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
