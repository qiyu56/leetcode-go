package problems

import (
	"container/heap"
	"sort"
)

// 田忌赛马

// https://leetcode-cn.com/problems/advantage-shuffle/
func AdvantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	maxHeap := &MaxHeap{}
	for i, e := range nums2 {
		heap.Push(maxHeap, Pair{i, e})
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	res := make([]int, n)
	left, right := 0, n-1
	for maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).(Pair)
		i, value := top.i, top.value
		if nums1[right] > value {
			res[i] = nums1[right]
			right--
		} else {
			res[i] = nums1[left]
			left++
		}
	}
	return res
}

type Pair struct {
	i     int
	value int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
