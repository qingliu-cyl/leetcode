package main

import "fmt"

func replaceSpace(s string) string {
	replaceStr := "%20"
	compareStr := " "
	res := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == compareStr {
			res += replaceStr
		} else {
			res += string(s[i])
		}
	}

	return res
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
