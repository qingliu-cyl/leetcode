package main

import "fmt"

func removeDigit(number string, digit byte) string {
	res := -1
	byteNumber := []byte(number)
	for idx, _ := range byteNumber {
		if digit != byteNumber[idx] {
			continue
		}

		if idx != len(byteNumber) - 1 {
			if byteNumber[idx+1] > byteNumber[idx] {
				return string(append(byteNumber[:idx], byteNumber[idx+1:]...))
			} else {
				res = idx
			}
		} else {
			return string(append(byteNumber[:idx]))
		}
	}
	if res == len(number) -1 {
		return string(byteNumber[:res])
	} else {
		return string(append(byteNumber[:res], byteNumber[res+1:]...))
	}
}

func main() {
	fmt.Println(removeDigit("123", []byte("3")[0]))
}
