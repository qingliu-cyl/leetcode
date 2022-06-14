package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	flag := []byte(" ")[0]

	str := ""
	list := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == flag && len(str) != 0 {
			list = append(list, str)
			str = ""
		} else if s[i] != flag {
			str += string(s[i])
		}
	}

	if len(str) != 0 {
		list = append(list, str)
	}

	p1, p2 := 0, len(list)-1
	for p1 < p2 {
		str1 := list[p1]
		list[p1] = list[p2]
		list[p2] = str1
		p1++
		p2--
	}

	return strings.Join(list, " ")
}

func main()  {
	fmt.Println(reverseWords("the sky is blue"))
}