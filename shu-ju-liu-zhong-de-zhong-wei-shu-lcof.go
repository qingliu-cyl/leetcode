package main

import "fmt"

type MedianFinder struct {
	nums []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}


func (this *MedianFinder) AddNum(num int)  {
	newNums := make([]int, len(this.nums)+1)
	s := 0
	add := false
	for i := 0; i < len(this.nums); s++{
		if this.nums[i] <= num || add {
			newNums[s] = this.nums[i]
			i++
		} else {
			newNums[s] = num
			add = true
		}
	}

	if !add {
		newNums[s] = num
	}

	this.nums = newNums

}


func (this *MedianFinder) FindMedian() float64 {
	if len(this.nums) == 0 {
		return 0
	}

	median := len(this.nums) / 2
	if len(this.nums) % 2 == 0 {
		return (float64(this.nums[median]) + float64(this.nums[median-1])) / float64(2)
	} else {
		return float64(this.nums[median])
	}
}

func main()  {
	m := Constructor()
	m.AddNum(-1)
	fmt.Println(m.FindMedian())
	m.AddNum(-1)
	fmt.Println(m.FindMedian())
	m.AddNum(-3)
}
