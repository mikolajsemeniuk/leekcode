package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/majority-element
//
// 169. Majority Element
//
// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// Example 1:
//
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
//
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 5 * 10^4
// -10^9 <= nums[i] <= 10^9
//
// Follow-up: Could you solve the problem in linear time and in O(1) space?
func majorityElement(nums []int) int {
	var candidate int
	var count int

	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
			continue
		}

		if num == candidate {
			count++
			continue
		}

		count--
	}

	return candidate
}

func TestMajorityElement(t *testing.T) {
	t.Run("nums = [3,2,3]", func(t *testing.T) {
		nums := []int{3, 2, 3}
		expected := 3
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [2,2,1,1,1,2,2]", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [6,5,5]", func(t *testing.T) {
		nums := []int{6, 5, 5}
		expected := 5
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,1,1,2,2]", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2}
		expected := 1
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [10,9,9,9,10]", func(t *testing.T) {
		nums := []int{10, 9, 9, 9, 10}
		expected := 9
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1,2,3,4,5,5,5,5,5]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 5, 5, 5, 5}
		expected := 5
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [8,8,7,7,7]", func(t *testing.T) {
		nums := []int{8, 8, 7, 7, 7}
		expected := 7
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [-1,1,1,1,2,1]", func(t *testing.T) {
		nums := []int{-1, 1, 1, 1, 2, 1}
		expected := 1
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [4,4,4,4,3,3,3]", func(t *testing.T) {
		nums := []int{4, 4, 4, 4, 3, 3, 3}
		expected := 4
		result := majorityElement(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
