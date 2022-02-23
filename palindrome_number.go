package main

import "fmt"

func isPalindrome(x int) bool {
	// 负数显然不是回文数
	if x < 0 {
		return false
	}
	xCopy := x
	res := 0
	for xCopy != 0 {
		res = res*10 + xCopy%10
		xCopy /= 10
	}

	return res == x
}

func main() {

	fmt.Println(isPalindrome(101))
}
