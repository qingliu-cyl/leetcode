package main

import "math"

type MinStack struct {
	stack, minStack []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt,
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if x < this.min {
		this.min = x
	}
	this.minStack = append(this.minStack, this.min)
}


func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}

	this.minStack = this.minStack[:len(this.minStack)-1]
	if len(this.minStack) == 0 {
		this.min = math.MaxInt
	} else {
		this.min = this.minStack[len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
	return this.min
}

func main() {

}
