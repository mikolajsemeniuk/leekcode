package leekcode

import "testing"

// https://leetcode.com/problems/search-insert-position/
//
// 35. Search Insert Position
//
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func TestSearchInsert(t *testing.T) {
	t.Run("target found in middle", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 5
		expected := 2
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("target not found, insert at position 1", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 2
		expected := 1
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("target not found, insert at end", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 7
		expected := 4
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("target not found, insert at beginning", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 0
		expected := 0
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("single element array, target found", func(t *testing.T) {
		nums := []int{1}
		target := 1
		expected := 0
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("single element array, target not found", func(t *testing.T) {
		nums := []int{1}
		target := 2
		expected := 1
		if result := searchInsert(nums, target); result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
