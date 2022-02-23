package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0 // 最长字符串的长度
	var strList map[uint8]struct{}
	for x, _ := range s {
		strList = make(map[uint8]struct{})
		for y := x; y < len(s); y++ {
			if _, ok := strList[s[y]]; ok {
				break
			} else {
				strList[s[y]] = struct{}{}
			}
		}
		if maxLen < len(strList) {
			maxLen = len(strList)
		}
	}

	if maxLen < len(strList) {
		return len(strList)
	} else {
		return maxLen
	}
}

func main() {
	lengthOfLongestSubstring("au")
}
