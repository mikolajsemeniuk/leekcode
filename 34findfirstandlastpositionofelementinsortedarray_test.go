package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
//
// 34. Find First and Last Position of Element in Sorted Array
//
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending
// position of a given target value.
//
// If target is not found in the array, return [-1, -1].
//
// Example 1:
//
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
//
// Example 2:
//
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
//
// Example 3:
//
// Input: nums = [], target = 0
// Output: [-1,-1]
//
// Constraints:
//
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
// nums is a non-decreasing array.
// -10^9 <= target <= 10^9
func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	for i, v := range nums {
		if v != target {
			continue
		}

		if start == -1 {
			start = i
		}

		end = i
	}

	return []int{start, end}
}

func TestSearchRange(t *testing.T) {
	t.Run("nums = [5,7,7,8,8,10], target = 8", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		expected := []int{3, 4}
		result := searchRange(nums, 8)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [5,7,7,8,8,10], target = 6", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		expected := []int{-1, -1}
		result := searchRange(nums, 6)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [], target = 0", func(t *testing.T) {
		nums := []int{}
		expected := []int{-1, -1}
		result := searchRange(nums, 0)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1], target = 1", func(t *testing.T) {
		nums := []int{1}
		expected := []int{0, 0}
		result := searchRange(nums, 1)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [2,2,2,2], target = 2", func(t *testing.T) {
		nums := []int{2, 2, 2, 2}
		expected := []int{0, 3}
		result := searchRange(nums, 2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
