package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/two-sum
//
// 1. Two Sum
//
// Given an array of integers nums and an integer target, return indices of the two numbers
// such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use
// the same element twice.
//
// You can return the answer in any order.
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
//
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
//
// Constraints:
//
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9
// Only one valid answer exists.
func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{i, j}
		}
		seen[n] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	t.Run("nums = [2,7,11,15], target = 9", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		expected := []int{0, 1}
		result := twoSum(nums, target)
		if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [3,2,4], target = 6", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		expected := []int{1, 2}
		result := twoSum(nums, target)
		if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [3,3], target = 6", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		expected := []int{0, 1}
		result := twoSum(nums, target)
		if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
