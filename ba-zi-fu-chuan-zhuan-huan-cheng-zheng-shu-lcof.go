package main

import "fmt"

func strToInt(str string) int {
	var min int64 = 0x100000000/2
	var max int64 = 0x100000000/2-1
	sign := 1
	var res int64 = 0
	start := false
	for idx, _ := range str {
		char := str[idx]

		if !start && []byte("-")[0] == char{
			start = true
			sign = -1
		} else if !start && []byte("+")[0] == char{
			start = true
			continue
		} else if !start && []byte(" ")[0] == char {
			continue
		} else if []byte("0")[0] <= char && char <= []byte("9")[0] {
			start = true
			res = res * 10 + int64(char - []byte("0")[0])
			if sign == 1 && res >= max {
				return  int(max)
			}

			if sign == -1 && res >= min{
				return -1 * int(min)
			}
		} else {
			break
		}
	}
	return int(res) * sign
}

func main()  {
	fmt.Println(strToInt(" -42"))
}