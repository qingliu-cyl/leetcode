package main

type MaxQueue struct {
	queue []int
	deque []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		deque: make([]int, 0),
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	for i := len(this.deque) -1; i >=0; i-- {
		if this.deque[i] > value {
			newQueue := make([]int, i+2)
			copy(newQueue, this.deque[:i+1])
			newQueue[i+1] = value
			this.deque = newQueue
			return
		} else if this.deque[i] == value {
			newQueue := make([]int, i+1)
			copy(newQueue, this.deque[:i+1])
			this.deque = newQueue
			return
		}
	}

	this.deque = []int{value}
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	res := this.queue[0]
	newQueue := make([]int, len(this.queue)-1)
	copy(newQueue, this.queue[1:])
	this.queue = newQueue

	if res  == this.deque[0] {
		newDeque := make([]int, len(this.deque)-1)
		copy(newDeque, this.deque[1:])
		this.deque = newDeque
	}

	return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {

}
