package main

func constructArr(a []int) []int {
	map1, map2 := make(map[int]int), make(map[int]int)
	map1[0] = 1
	for i := 1; i < len(a); i++ {
		map1[i] = map1[i-1] * a[i-1]
	}

	map2[len(a)-1] = 1
	for i := len(a)-2; i >= 0; i-- {
		map2[i] = map2[i+1] * a[i+1]
	}

	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = map1[i] * map2[i]
	}
	return res
}

func main()  {
	constructArr([]int{1,2,3,4,5})
}
