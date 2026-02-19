package leekcode

import (
	"math/rand"
	"testing"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array
//
// 215. Kth Largest Element in an Array
//
// Given an integer array nums and an integer k, return the kth largest element in the array.
//
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Example 1:
//
// Input: nums = [3,2,1,5,6,4], k = 2
// Output: 5
//
// Example 2:
//
// Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
// Output: 4
//
// Constraints:
//
// 1 <= k <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
func findKthLargest(a []int, k int) int {
	n := len(a)
	l := 0
	r := n - 1
	for l < r {
		i, j := partition(a, l, r)
		if n-k <= i {
			r = i
		} else if n-k > j {
			l = j + 1
		} else {
			return a[n-k]
		}
	}
	return a[n-k]
}

func partition(a []int, l, r int) (int, int) {
	pivot := a[l+rand.Intn(r-l+1)]
	i := l
	for i <= r {
		if a[i] < pivot {
			a[i], a[l] = a[l], a[i]
			l++
			i++
		} else if a[i] > pivot {
			a[i], a[r] = a[r], a[i]
			r--
		} else {
			i++
		}
	}
	return l - 1, r
}

func TestFindKthLargest(t *testing.T) {
	t.Run("nums = [3,2,1,5,6,4], k = 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 5, 6, 4}
		expected := 5
		result := findKthLargest(nums, 2)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [3,2,3,1,2,4,5,5,6], k = 4", func(t *testing.T) {
		nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
		expected := 4
		result := findKthLargest(nums, 4)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1], k = 1", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := findKthLargest(nums, 1)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [2,1], k = 2", func(t *testing.T) {
		nums := []int{2, 1}
		expected := 1
		result := findKthLargest(nums, 2)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [7,6,5,4,3,2,1], k = 1", func(t *testing.T) {
		nums := []int{7, 6, 5, 4, 3, 2, 1}
		expected := 7
		result := findKthLargest(nums, 1)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [5,5,5,5], k = 3", func(t *testing.T) {
		nums := []int{5, 5, 5, 5}
		expected := 5
		result := findKthLargest(nums, 3)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-1,-1,0,2,4], k = 4", func(t *testing.T) {
		nums := []int{-1, -1, 0, 2, 4}
		expected := -1
		result := findKthLargest(nums, 4)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
