package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/move-zeroes
//
// 283. Move Zeroes
//
// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
//
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
//
// Input: nums = [0]
// Output: [0]
//
// Constraints:
//
// 1 <= nums.length <= 10^4
// -2^31 <= nums[i] <= 2^31 - 1
//
// Follow up: Could you minimize the total number of operations done?
func moveZeroes(nums []int) {
	// Two-pointer approach
	// leftPtr points to the position where the next non-zero should be placed
	leftPtr := 0

	// Move all non-zero elements to the front
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[leftPtr] = nums[i]
			leftPtr++
		}
	}

	// Fill the rest with zeros
	for i := leftPtr; i < len(nums); i++ {
		nums[i] = 0
	}
}

func TestMoveZeroes(t *testing.T) {
	t.Run("nums = [0,1,0,3,12]", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 12}
		expected := []int{1, 3, 12, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [0]", func(t *testing.T) {
		nums := []int{0}
		expected := []int{0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1,2,3,4,5]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [0,0,0,0,0]", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 0}
		expected := []int{0, 0, 0, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1,0,2,0,3,0,4]", func(t *testing.T) {
		nums := []int{1, 0, 2, 0, 3, 0, 4}
		expected := []int{1, 2, 3, 4, 0, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [0,0,1]", func(t *testing.T) {
		nums := []int{0, 0, 1}
		expected := []int{1, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [2,1]", func(t *testing.T) {
		nums := []int{2, 1}
		expected := []int{2, 1}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [4,2,4,0,0,3,0,5,1,0]", func(t *testing.T) {
		nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
		expected := []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := []int{1}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})

	t.Run("nums = [-1,0,2,0,-3]", func(t *testing.T) {
		nums := []int{-1, 0, 2, 0, -3}
		expected := []int{-1, 2, -3, 0, 0}
		moveZeroes(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, got %v", expected, nums)
		}
	})
}
