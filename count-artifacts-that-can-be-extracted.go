package main

import "strconv"

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	artifactAddr := func(artifact []int) [][]int {
		res := make([][]int, 0)
		for i := artifact[0]; i <= artifact[2]; i++ {
			for j := artifact[1]; j <= artifact[3]; j++ {
				res = append(res, []int{i, j})
			}
		}

		return res
	}

	// 可挖掘的位置
	digMap := make(map[string]struct{})
	for _, v := range dig {
		key := strconv.Itoa(v[0]) + "-" + strconv.Itoa(v[1])
		digMap[key] = struct{}{}
	}
	res := 0
	for _, artifact := range artifacts {
		addrs := artifactAddr(artifact)
		extracted := true
		for _, addr := range addrs {
			key := strconv.Itoa(addr[0]) + "-" + strconv.Itoa(addr[1])
			if _, ok := digMap[key]; !ok {
				extracted = false
				break
			}
		}

		if extracted {
			res++
		}
	}

	return res
}

func main() {

}
