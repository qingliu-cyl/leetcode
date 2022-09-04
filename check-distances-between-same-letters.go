package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	aByte := []byte("a")[0]

	for i := 0; i < len(s); i++ {
		dist := distance[s[i]-aByte]
		if  i+dist+1 >= len(s) {
			return false
		}
		if s[i] == s[i+dist+1] {
			distance[s[i]-aByte] = -1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkDistances("aa", []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
}
