package main

import "fmt"

func countCollisions(directions string) int {
	//以下三种情形会相撞
	collisionsMap := make(map[byte][]byte)
	collisionsMap[[]byte("R")[0]] = []byte("L, S")
	collisionsMap[[]byte("S")[0]] = []byte("L")
	RCarNum := 0

	list := []byte(directions)
	judge := func(cur, next int) int {
		if _, ok := collisionsMap[list[cur]]; !ok {
			return 0
		}

		for _, v := range collisionsMap[list[cur]] {
			if v == list[next] {
				if string(v) == "S" || string(list[cur]) == "S" {
					return 1
				} else {
					return 2
				}
			}
		}

		return 0
	}

	num := 0
	for i := 0; i < len(list)-1; i++ {
		add := judge(i, i+1)
		// 碰撞之后车停下来
		if add > 0 {
			num += add
			num += RCarNum
			RCarNum = 0
			list[i], list[i+1] = []byte("S")[0], []byte("S")[0]
		} else {
			if list[i] == []byte("R")[0] {
				RCarNum++
			}
		}
	}
	return num
}
func main() {
	fmt.Println(countCollisions("RLRSLL"))
}
