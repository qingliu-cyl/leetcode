package main

import "fmt"

func largestGoodInteger(num string) string {
	if len(num) < 3 {
		return ""
	}

	max := func(i, j byte) byte {
		if i > j {
			return i
		} else {
			return j
		}
	}
	has := false
	var res byte
	q := 2
	for ; q < len(num); q++ {
		if num[q] == num[q-1] && num[q] == num[q-2] {
			if !has {
				res = num[q]
			} else {
				res = max(res, num[q])
			}
			has = true
		}
	}

	if has {
		return string([]byte{res, res, res})
	} else {
		return ""
	}
}

func main() {
	fmt.Println(largestGoodInteger("2300019"))
}
