package main

import "fmt"

// MaxSlidingWindow returns the maximum of every length-k window. A deque of
// indexes holds a decreasing front-runner list: the front is always the
// current window's maximum.
func MaxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return []int{}
	}
	if k > len(nums) {
		k = len(nums)
	}
	dq := []int{} // indexes; their values are strictly decreasing
	out := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:] // front fell out of the window
		}
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= v {
			dq = dq[:len(dq)-1] // dominated: smaller and older than v
		}
		dq = append(dq, i)
		if i >= k-1 {
			out = append(out, nums[dq[0]])
		}
	}
	return out
}

func main() {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
