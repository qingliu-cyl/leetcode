package main

import "fmt"

// 超时了
//https://leetcode-cn.com/contest/weekly-contest-285/problems/longest-substring-of-one-repeating-character/
//https://leetcode-cn.com/contest/weekly-contest-285
func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	k := len(queryIndices)
	res := make([]int, k)
	max := 1
	maxL := 0
	maxR := 0

	// 重复长度
	l, r, repeatLen := 0, 0, 1
	sByte := []byte(s)
	// 第一次查询时的最大值
	getMax := func() {
		for i := 1; i < len(s); i++ {
			if sByte[i] == sByte[i-1] {
				r = i
				repeatLen++
				if max < repeatLen {
					maxL = l
					maxR = r
					max = repeatLen
				}
			} else {
				if max < repeatLen {
					maxL = l
					maxR = r
					max = repeatLen
				}
				l, r, repeatLen = i, i, 1
			}
		}
	}

	getMax()

	judgeLen := func(idx int) (int, int, int) {
		res := 1
		i := idx - 1
		for i >= 0 && sByte[i] == sByte[idx] {
			res++
			i--
		}
		l := i + 1

		i = idx + 1
		for i < len(sByte) && sByte[i] == sByte[idx] {
			i++
			res++
		}
		r := i - 1

		return l, r, res
	}

	for idx, i := range queryIndices {
		sByte[i] = queryCharacters[idx]
		if maxL <= i && maxR >= i {
			max = 1
			maxL = 0
			maxR = 0
			repeatLen = 1
			getMax()
		} else {
			l, r, nowLen := judgeLen(i)
			if nowLen > max {
				max = nowLen
				maxL = l
				maxR = r
			}
		}
		res[idx] = max
	}

	return res
}

func main() {
	//	"geuqjmt"
	//	"bgemoegklm"
	//[3,4,2,6,5,6,5,4,3,2]
	fmt.Println(longestRepeating("geuqjmt", "bgemoegklm", []int{3, 4, 2, 6, 5, 6, 5, 4, 3, 2}))
}
