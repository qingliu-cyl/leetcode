package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			total++
		}
	}

	return total * 100 / len(s)
}

func main() {
	fmt.Println(percentageLetter("vmvvvvvzrvvpvdvvvvyfvdvvvvpkvvbvvkvvfkvvvkvbvvnvvomvzvvvdvvvkvvvvvvvvvlvcvilaqvvhoevvlmvhvkvtgwfvvzy", []byte("v")[0]))
}
