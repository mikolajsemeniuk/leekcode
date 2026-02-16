package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/jump-game
//
// 55. Jump Game
//
// You are given an integer array nums. You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//
// Example 2:
//
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0,
// which makes it impossible to reach the last index.
//
// Constraints:
//
// 1 <= nums.length <= 10^4
// 0 <= nums[i] <= 10^5
func canJump(nums []int) bool {
	maxReach := 0
	lastIndex := len(nums) - 1

	for i := 0; i <= maxReach; i++ {
		val := nums[i]
		if i+val > maxReach {
			maxReach = i + val
		}

		if maxReach >= lastIndex {
			return true
		}
	}

	return false
}

func TestCanJump(t *testing.T) {
	t.Run("example 2: [1,0,1,0,4]", func(t *testing.T) {
		nums := []int{1, 0, 1, 0, 4}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 1: [2,3,1,1,4]", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 2: [3,2,1,0,4]", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		nums := []int{0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("two elements - can reach", func(t *testing.T) {
		nums := []int{1, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("two elements - cannot reach", func(t *testing.T) {
		nums := []int{0, 1}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("all zeros except first", func(t *testing.T) {
		nums := []int{2, 0, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zero in middle blocks path", func(t *testing.T) {
		nums := []int{1, 0, 1, 0}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("large jump at start", func(t *testing.T) {
		nums := []int{5, 0, 0, 0, 0, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("can skip zeros", func(t *testing.T) {
		nums := []int{3, 0, 0, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("minimum jumps needed", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("one big jump", func(t *testing.T) {
		nums := []int{10, 0, 0, 0, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("cannot reach end - zero trap", func(t *testing.T) {
		nums := []int{1, 1, 0, 1}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("exactly reach end", func(t *testing.T) {
		nums := []int{2, 0, 0}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("just short of end", func(t *testing.T) {
		nums := []int{1, 0, 0}
		expected := false
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("alternating high low", func(t *testing.T) {
		nums := []int{5, 1, 5, 1, 5}
		expected := true
		result := canJump(nums)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
