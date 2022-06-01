package main

type CQueue struct {
	que []int
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.que = append(this.que, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.que) == 0 {
		return -1
	}

	value := this.que[0]
	if len(this.que) == 1 {
		this.que = nil
	} else {
		this.que = this.que[1:]
	}
	return value
}

func main() {

}
