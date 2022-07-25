package main

import "fmt"

func permutation(s string) []string {
	list := []byte(s)
	res := make([]string, 0)
	var dfs func(idx int)
	dfs  = func(idx int) {
		if idx == len(s)-1 {
			res = append(res, string(list))
			return
		}

		dic := make(map[byte]bool)
		for i := idx; i < len(list); i++ {
			if dic[list[i]] {
				continue
			}

			dic[list[i]] = true

			head := list[idx]
			list[idx] = list[i]
			list[i] = head
			dfs(idx+1)
			head = list[idx]
			list[idx] = list[i]
			list[i] = head
		}
	}

	dfs(0)
	return res
}

func main() {
	fmt.Println(permutation("abb"))
}
