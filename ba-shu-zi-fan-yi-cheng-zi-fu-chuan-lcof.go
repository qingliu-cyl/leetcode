package main

func translateNum(num int) int {
	if num == 0 {
		return 1
	}

	res := make([]int, 0)
	for num != 0 {
		res = append(res, num % 10)
		num = num / 10
	}

	count := make([]int, len(res)+1)
	count[len(res)], count[len(res)-1] = 1, 1
	for i := len(res)-2; i >= 0; i-- {
		if res[i+1] != 0 && res[i+1] * 10 + res[i] <= 25 {
			count[i] = count[i+1] + count[i+2]
		} else {
			count[i] = count[i+1]
		}
	}

	return count[0]
}