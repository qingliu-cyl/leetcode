package main

import "fmt"

func longestPalindrome(s string) string {
	data := []byte(s)
	dataS := 0
	dataE := 0
	doFunc := func(start, end int) {
		for {
			if start-1 < 0 || end+1 >= len(data) || data[start-1] != data[end+1] {
				if end-start > dataE-dataS {
					dataE = end
					dataS = start
				}
				return
			} else {
				start--
				end++
			}
		}

	}

	for i := 0; i < len(data); i++ {
		if i+1 < len(data) && i-1 >= 0 && data[i+1] == data[i-1] {
			doFunc(i-1, i+1)
		}
		if i+1 < len(data) && data[i] == data[i+1] {
			doFunc(i, i+1)
		}
		if i-1 >= 0 && data[i] == data[i-1] {
			doFunc(i-1, i)
		}
	}
	return s[dataS : dataE+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
