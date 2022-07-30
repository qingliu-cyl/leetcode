package main

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}

	push := func(idx int) {
		for len(queue) > 0 && nums[idx] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, idx)
	}

	for i := 0; i < k; i++ {
		push(i)
	}


	res = append(res, nums[queue[0]])

	for i := k; i < len(nums); i++ {
		push(i)
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}

	return res
}

func main() {

}
