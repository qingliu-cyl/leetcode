package main

import "fmt"

func convert(s string, numRows int) string {
	res := make([][]byte, numRows)
	add := true
	i := 0 // 用i控制放入第几行
	// 行数为1时直接输出
	if numRows == 1 {
		return s
	}
	// 遍历
	for idx, b := range []byte(s) {
		res[i] = append(res[i], b)
		if (i == numRows-1 || i == 0) && idx != 0 { // 当到达边界时折返
			add = !add
		}
		if add {
			i++
		} else {
			i--
		}
	}

	// 输出
	ret := ""
	for _, i := range res {
		ret += string(i)
	}

	return ret

}

func main() {
	fmt.Println(convert("AB", 1))
}
