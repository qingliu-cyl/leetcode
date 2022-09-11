package main

import "fmt"

func partitionString(s string) int {
	ret := make(map[byte]struct{})

	num := 0
	for i,_  := range s {
		if _, ok := ret[s[i]]; ok {
			ret = make(map[byte]struct{})
			num++
		}
		ret[s[i]] = struct{}{}
	}

	if len(ret) != 0 {
		num++
	}

	return num
}

func main()  {
	fmt.Println(partitionString("sssssss"))
}