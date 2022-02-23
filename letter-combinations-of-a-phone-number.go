package main

import "fmt"

var res []string

var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func backtrack(digits string, deep int, str string) {
	if deep == len(digits) {
		if len(str) == 0 {
			return
		}
		res = append(res, str)
	} else {
		for _, i := range dict[string(digits[deep])] {
			backtrack(digits, deep+1, str+i)
		}
	}
}

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	backtrack(digits, 0, "")
	return res
}

func main() {
	fmt.Println(letterCombinations(""))
}
