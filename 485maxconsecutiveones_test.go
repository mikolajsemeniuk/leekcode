package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/max-consecutive-ones
//
// 485. Max Consecutive Ones
//
// Given a binary array nums, return the maximum number of consecutive 1's in the array.
//
// Example 1:
//
// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s.
// The maximum number of consecutive 1s is 3.
//
// Example 2:
//
// Input: nums = [1,0,1,1,0,1]
// Output: 2
//
// Constraints:
//
// 1 <= nums.length <= 10^5
// nums[i] is either 0 or 1.
func findMaxConsecutiveOnes(nums []int) int {
	var current int
	var total int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
			if current > total {
				total = current
			}
		}

		if nums[i] != 1 {
			current = 0
		}
	}

	return total
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	t.Run("nums = [1,1,0,1,1,1]", func(t *testing.T) {
		nums := []int{1, 1, 0, 1, 1, 1}
		expected := 3
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,0,1,1,0,1]", func(t *testing.T) {
		nums := []int{1, 0, 1, 1, 0, 1}
		expected := 2
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0]", func(t *testing.T) {
		nums := []int{0}
		expected := 0
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,1,1,1,1]", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1}
		expected := 5
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,0,0,0,0]", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 0}
		expected := 0
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,0,1,1,1,0,1,1,1,1,0,1]", func(t *testing.T) {
		nums := []int{1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1}
		expected := 4
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,1,1,0,1,1,1]", func(t *testing.T) {
		nums := []int{0, 1, 1, 0, 1, 1, 1}
		expected := 3
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,1,0,0,1,1]", func(t *testing.T) {
		nums := []int{1, 1, 0, 0, 1, 1}
		expected := 2
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,0,1,0,1,0,1,0]", func(t *testing.T) {
		nums := []int{1, 0, 1, 0, 1, 0, 1, 0}
		expected := 1
		result := findMaxConsecutiveOnes(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
