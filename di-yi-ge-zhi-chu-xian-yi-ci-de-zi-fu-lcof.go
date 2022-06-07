package main

func firstUniqChar(s string) byte {
	resMap := make(map[int32]int)
	for _, v := range s {
		if _, ok := resMap[v]; ok {
			resMap[v]++
		} else {
			resMap[v] = 1
		}
	}

	for _, v := range s {
		if resMap[v] == 1 {
			return byte(v)
		}
	}

	return []byte(" ")[0]
}

func main() {

}
