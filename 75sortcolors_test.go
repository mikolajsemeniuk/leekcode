package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/sort-colors
//
// 75. Sort Colors
//
// Given an array nums with n objects colored red, white, or blue, sort them in-place
// so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue.
//
// You must solve this problem without using the library's sort function.
//
// Example 1:
//
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
//
// Example 2:
//
// Input: nums = [2,0,1]
// Output: [0,1,2]
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.
//
// Follow up: Could you come up with a one-pass algorithm using only constant extra space?
func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func TestSortColors(t *testing.T) {
	t.Run("nums = [2,0,2,1,1,0]", func(t *testing.T) {
		nums := []int{2, 0, 2, 1, 1, 0}
		expected := []int{0, 0, 1, 1, 2, 2}
		sortColors(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [2,0,1]", func(t *testing.T) {
		nums := []int{2, 0, 1}
		expected := []int{0, 1, 2}
		sortColors(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [0]", func(t *testing.T) {
		nums := []int{0}
		expected := []int{0}
		sortColors(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1,1,1]", func(t *testing.T) {
		nums := []int{1, 1, 1}
		expected := []int{1, 1, 1}
		sortColors(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [2,2,0,1,0,2,1,0]", func(t *testing.T) {
		nums := []int{2, 2, 0, 1, 0, 2, 1, 0}
		expected := []int{0, 0, 0, 1, 1, 2, 2, 2}
		sortColors(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})
}
