package main


func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}

	l,r,t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++

		if  t  > b {
			break
		}
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--

		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--

		if b < t {
			break
		}
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if r < l {
			break
		}
	}

	return res
}

func main()  {
	spiralOrder([][]int{{1}})
}