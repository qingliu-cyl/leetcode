package main


func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, len(pushed))
	index := 0
	top := -1
	for i := 0; i < len(pushed); i++ {
		top++
		stack[top] = pushed[i]
		for  top >= 0 && stack[top] == popped[index] {
			// 出栈
			index++
			top--
		}
	}

	return top == -1
}
