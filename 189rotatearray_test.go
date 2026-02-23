package leekcode

import (
	"reflect"
	"slices"
	"testing"
)

// https://leetcode.com/problems/rotate-array
//
// 189. Rotate Array
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
//
// Example 2:
//
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
//
// Constraints:
//
// 1 <= nums.length <= 10^5
// -2^31 <= nums[i] <= 2^31 - 1
// 0 <= k <= 10^5
//
// Follow up:
// Try to come up with as many solutions as you can. There are at least three different ways
// to solve this problem. Could you do it in-place with O(1) extra space?
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k %= n
	if k == 0 {
		return
	}

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func TestRotate(t *testing.T) {
	t.Run("nums = [1,2,3,4,5,6,7], k = 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		rotate(nums, 3)
		expected := []int{5, 6, 7, 1, 2, 3, 4}
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [-1,-100,3,99], k = 2", func(t *testing.T) {
		nums := []int{-1, -100, 3, 99}
		rotate(nums, 2)
		expected := []int{3, 99, -1, -100}
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1], k = 10", func(t *testing.T) {
		nums := []int{1}
		rotate(nums, 10)
		expected := []int{1}
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1,2], k = 1", func(t *testing.T) {
		nums := []int{1, 2}
		rotate(nums, 1)
		expected := []int{2, 1}
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})
}
