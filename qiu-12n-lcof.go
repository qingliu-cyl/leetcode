package main

func sumNums(n int) int {
	res := 0
	var add func(n int) bool
	add = func(n int) bool {
		res += n
		return n > 0 && add(n-1)
	}

	add(n)
	return res
}
