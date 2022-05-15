package main

import "fmt"

func removeAnagrams(words []string) []string {
	l := len(words)
	compare := func(str1, str2 string) bool {
		strMap := make(map[int32]int)
		for _, i := range str1 {
			strMap[i] += 1
		}
		for _, j := range str2 {
			if strMap[j] == 0 {
				return false
			}
			strMap[j]--
		}
		return true
	}

	for i := 1; i < l; {
		if len(words[i]) != len(words[i-1]) {
			i++
			continue
		}

		if compare(words[i], words[i-1]) {
			words = append(words[:i], words[i+1:]...)
			l = len(words)
		} else {
			i++
		}
	}

	return words
}

func main() {
	fmt.Println(removeAnagrams([]string{"abba","baba","bbaa","cd","cd"}))
}
