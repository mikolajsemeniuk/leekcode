package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/maximum-subarray
//
// 53. Maximum Subarray
//
// Given an integer array nums, find the subarray with the largest sum, and return its sum.
//
// Example 1:
//
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
//
// Example 2:
//
// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.
//
// Example 3:
//
// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
//
// Constraints:
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
//
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
func maxSubArray(nums []int) int {
	// Kadane's Algorithm
	// The idea is to keep track of the maximum sum ending at the current position
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either extend the existing subarray or start a new one
		val := nums[i]
		currentSum = max(val, currentSum+val)
		// Update the maximum sum found so far
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func TestMaxSubArray(t *testing.T) {
	t.Run("nums = [-2,1,-3,4,-1,2,1,-5,4]", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		expected := 6
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [5,4,-1,7,8]", func(t *testing.T) {
		nums := []int{5, 4, -1, 7, 8}
		expected := 23
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-1]", func(t *testing.T) {
		nums := []int{-1}
		expected := -1
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-2,-1]", func(t *testing.T) {
		nums := []int{-2, -1}
		expected := -1
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,2,3,4,5]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 15
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-2,1,-3,4,-1,2,1,-5,4]", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		expected := 6
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [8,-19,5,-4,20]", func(t *testing.T) {
		nums := []int{8, -19, 5, -4, 20}
		expected := 21
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-1,-2,-3,-4]", func(t *testing.T) {
		nums := []int{-1, -2, -3, -4}
		expected := -1
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [2,1,-3,4]", func(t *testing.T) {
		nums := []int{2, 1, -3, 4}
		expected := 4
		result := maxSubArray(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
