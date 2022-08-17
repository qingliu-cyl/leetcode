package main

import (
	"fmt"
)

func lastRemaining(n int, m int) int {
	f := 0
	for i := 2; i != n+1; i++ {
		f =  (f + m) % i
	}

	return f
}

func main() {
	fmt.Println(lastRemaining(5, 3))
}
