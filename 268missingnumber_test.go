package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/missing-number
//
// 268. Missing Number
//
// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
//
// Example 1:
//
// Input: nums = [3,0,1]
// Output: 2
// Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
// 2 is the missing number in the range since it does not appear in nums.
//
// Example 2:
//
// Input: nums = [0,1]
// Output: 2
// Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
// 2 is the missing number in the range since it does not appear in nums.
//
// Example 3:
//
// Input: nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
// Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
// 8 is the missing number in the range since it does not appear in nums.
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 10^4
// 0 <= nums[i] <= n
// All the numbers of nums are unique.
//
// Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
func missingNumber(nums []int) int {
	// Using sum formula: sum of 0 to n = n * (n + 1) / 2
	// Missing number = expected sum - actual sum
	n := len(nums)
	expected := (n + 1) * n / 2
	actual := 0

	for _, num := range nums {
		actual += num
	}

	return expected - actual
}

func TestMissingNumber(t *testing.T) {
	t.Run("nums = [3,0,1]", func(t *testing.T) {
		nums := []int{3, 0, 1}
		expected := 2
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,1]", func(t *testing.T) {
		nums := []int{0, 1}
		expected := 2
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [9,6,4,2,3,5,7,0,1]", func(t *testing.T) {
		nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		expected := 8
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0]", func(t *testing.T) {
		nums := []int{0}
		expected := 1
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 0
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,1,2,3,4,5,6,7,9]", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
		expected := 8
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,2,3,4,5,6,7,8,9]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := 0
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,1,2,3,4,5,6,7,8]", func(t *testing.T) {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		expected := 9
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [2,0]", func(t *testing.T) {
		nums := []int{2, 0}
		expected := 1
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [45,35,38,13,12,23,48,29,25,33,19,0,42,17,1,18,39,22,49,37,30,24,10,5,47,26,14,2,31,21,15,46,20,27,3,11,6,40,28,36,9,32,8,16,43,41,7,4,34]", func(t *testing.T) {
		nums := []int{45, 35, 38, 13, 12, 23, 48, 29, 25, 33, 19, 0, 42, 17, 1, 18, 39, 22, 49, 37, 30, 24, 10, 5, 47, 26, 14, 2, 31, 21, 15, 46, 20, 27, 3, 11, 6, 40, 28, 36, 9, 32, 8, 16, 43, 41, 7, 4, 34}
		expected := 44
		result := missingNumber(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
