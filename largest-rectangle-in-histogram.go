package main

import "fmt"

type Stack struct {
	stack []int
	top   int
}

// 弹出栈顶
func (s *Stack) pop() {
	if s.top == -1 {
		return
	} else {
		s.top--
	}
}

// 获取栈顶的值
func (s *Stack) getTopValue() int {
	return s.stack[s.top]
}

// 获取栈顶元素左侧的值
func (s *Stack) getTopLeftValue() int {
	if s.top == 0 {
		return -1
	} else {
		return s.stack[s.top-1]
	}
}

// 栈是否为空
func (s *Stack) isEmpty() bool {
	return s.top == -1
}

// 像栈顶添加一个元素
func (s *Stack) append(v int) {
	s.top++
	s.stack[s.top] = v
}

func largestRectangleArea(heights []int) int {
	// 初始化栈
	res := 0
	stack := &Stack{
		stack: make([]int, len(heights)),
		top:   -1,
	}

	setRes := func(i int) {
		if i > res {
			res = i
		}
	}

	for idx, v := range heights {
		for true {
			if stack.isEmpty() {
				stack.append(idx)
				break
			} else if heights[stack.getTopValue()] > v {
				// 此时应该计算面积并出栈
				area := (idx - (stack.getTopLeftValue()) - 1) * heights[stack.getTopValue()]
				setRes(area)
				stack.pop()
			} else {
				stack.append(idx)
				break
			}
		}
	}

	for !stack.isEmpty() {
		area := (len(heights) - (stack.getTopLeftValue()) - 1) * heights[stack.getTopValue()]
		setRes(area)
		stack.pop()
	}

	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}
