package leekcode

import (
	"container/heap"
	"testing"
)

// https://leetcode.com/problems/find-median-from-data-stream
//
// 295. Find Median from Data Stream
//
// The median is the middle value in an ordered integer list. If the size of the list is even,
// the median is the average of the two middle values.
//
// Implement the MedianFinder class:
// - MedianFinder() initializes the MedianFinder object.
// - void addNum(int num) adds the integer num from the data stream to the data structure.
// - double findMedian() returns the median of all elements so far.
//
// Example 1:
//
// Input
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// Output
// [null, null, null, 1.5, null, 2.0]
//
// Constraints:
//
// -10^5 <= num <= 10^5
// There will be at least one element in the data structure before calling findMedian.
// At most 5 * 10^4 calls will be made to addNum and findMedian.
type MedianFinder struct {
	lower *maxHeap
	upper *minHeap
}

func Constructor() MedianFinder {
	lower := &maxHeap{}
	upper := &minHeap{}
	heap.Init(lower)
	heap.Init(upper)
	return MedianFinder{lower: lower, upper: upper}
}

func (m *MedianFinder) AddNum(num int) {
	if m.lower.Len() == 0 || num <= (*m.lower)[0] {
		heap.Push(m.lower, num)
	} else {
		heap.Push(m.upper, num)
	}

	if m.lower.Len() > m.upper.Len()+1 {
		heap.Push(m.upper, heap.Pop(m.lower))
	} else if m.upper.Len() > m.lower.Len() {
		heap.Push(m.lower, heap.Pop(m.upper))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.lower.Len() > m.upper.Len() {
		return float64((*m.lower)[0])
	}
	left := (*m.lower)[0]
	right := (*m.upper)[0]
	return float64(left+right) / 2.0
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func TestMedianFinder(t *testing.T) {
	t.Run("example sequence", func(t *testing.T) {
		medianFinder := Constructor()
		medianFinder.AddNum(1)
		medianFinder.AddNum(2)
		if got := medianFinder.FindMedian(); got != 1.5 {
			t.Errorf("Expected 1.5, got %v", got)
		}
		medianFinder.AddNum(3)
		if got := medianFinder.FindMedian(); got != 2.0 {
			t.Errorf("Expected 2.0, got %v", got)
		}
	})

	t.Run("even count", func(t *testing.T) {
		medianFinder := Constructor()
		medianFinder.AddNum(5)
		medianFinder.AddNum(3)
		medianFinder.AddNum(8)
		medianFinder.AddNum(9)
		if got := medianFinder.FindMedian(); got != 6.5 {
			t.Errorf("Expected 6.5, got %v", got)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		medianFinder := Constructor()
		medianFinder.AddNum(-1)
		medianFinder.AddNum(-2)
		medianFinder.AddNum(-3)
		if got := medianFinder.FindMedian(); got != -2.0 {
			t.Errorf("Expected -2.0, got %v", got)
		}
	})
}
