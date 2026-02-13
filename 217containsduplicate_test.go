package leekcode

import (
	"sort"
	"testing"
)

// https://leetcode.com/problems/contains-duplicate
//
// 217. Contains Duplicate
//
// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
//
// Example 1:
//
// Input: nums = [1,2,3,1]
// Output: true
//
// Example 2:
//
// Input: nums = [1,2,3,4]
// Output: false
//
// Example 3:
//
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true
//
// Constraints:
//
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func containsDuplicate(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	t.Run("nums = [1,2,3,1]", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1,2,3,4]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := false
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1,1,1,3,3,4,3,2,4,2]", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := false
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1,2]", func(t *testing.T) {
		nums := []int{1, 2}
		expected := false
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1,1]", func(t *testing.T) {
		nums := []int{1, 1}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [1,5,-2,-4,0]", func(t *testing.T) {
		nums := []int{1, 5, -2, -4, 0}
		expected := false
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [3,3]", func(t *testing.T) {
		nums := []int{3, 3}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [0,4,5,0,3,6]", func(t *testing.T) {
		nums := []int{0, 4, 5, 0, 3, 6}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [2,14,18,22,22]", func(t *testing.T) {
		nums := []int{2, 14, 18, 22, 22}
		expected := true
		result := containsDuplicate(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
