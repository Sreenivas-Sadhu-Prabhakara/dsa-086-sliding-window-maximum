package main

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{9, 8, 7}, 2, []int{9, 8}},
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{4, 4, 4}, 2, []int{4, 4}},
		{[]int{1, 2}, 5, []int{2}}, // k larger than input: one whole-array window
		{[]int{}, 3, []int{}},
	}
	for _, c := range cases {
		if got := MaxSlidingWindow(c.in, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxSlidingWindow(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
