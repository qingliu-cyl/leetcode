package main

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	
	board := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
	}

	canMove := func(i, j int) bool {
		num := 0
		for i != 0 {
			num += i%10
			i /= 10
		}

		for j != 0 {
			num += j%10
			j /= 10
		}
		if num > k {
			return false
		} else {
			return true
		}
	}

	type Key struct {
		i int
		j int
	}
	moveMap := make(map[Key]struct{})
	
	var move func(i, j int)
	move = func(i, j int) {
		if i >= m || j >= n || i < 0 || j < 0 || !canMove(i, j){
			return
		}

		if _, ok := moveMap[Key{i,j}]; ok {
			return
		}

		moveMap[Key{i,j}] = struct{}{}
		move(i+1,j)
		move(i-1,j)
		move(i,j+1)
		move(i,j-1)
	}

	move(0,0)

	return len(moveMap)
}