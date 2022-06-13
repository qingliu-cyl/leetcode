package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	index := func(list []byte, target byte) int {
		for idx, i := range list {
			if i == target {
				return idx
			}
		}

		return -1
	}

	res := make([][]byte, len(s))
	res[0] = []byte{s[0]}
	max := 1
	for i := 1 ; i < len(s); i++ {
		idx := index(res[i-1], s[i])

		if idx == -1 {
			res[i] = append(res[i-1], s[i])
		} else {
			if idx == len(res[i-1]) -1 {
				res[i] = []byte{s[i]}
			} else {
				res[i] = append(res[i-1][idx+1:], s[i])
			}
		}
		if len(res[i]) > max {
			max = len(res[i])
		}
	}

	return max
}