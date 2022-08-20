package main

func verifyPostorder(postorder []int) bool {
	var  reduce func(i, j int) bool

	reduce = func(i, j int) bool {
		if i >= j {
			return true
		}

		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		m := p
		for postorder[p] > postorder[j] {
			p++
		}

		return p == j && reduce(i, m-1) && reduce(m, j-1)
	}

	return reduce(0, len(postorder)-1)
}
