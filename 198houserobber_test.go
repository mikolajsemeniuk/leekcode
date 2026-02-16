package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/house-robber
//
// 198. House Robber
//
// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
// the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and
// it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can
// rob tonight without alerting the police.
//
// Example 1:
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
//
// Example 2:
//
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.
//
// Constraints:
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// dp[i] represents the maximum amount we can rob up to house i
	// At each house, we have two choices:
	// 1. Rob this house: nums[i] + dp[i-2] (can't rob previous house)
	// 2. Don't rob this house: dp[i-1] (take the max from previous house)

	// Space optimized: we only need to track the last two values
	twoHousesBack := nums[0]
	oneHouseBack := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		val := nums[i]
		current := max(val+twoHousesBack, oneHouseBack)
		twoHousesBack = oneHouseBack
		oneHouseBack = current
	}

	return oneHouseBack
}

func TestRob(t *testing.T) {
	t.Run("example 1: [100,1,2,100]", func(t *testing.T) {
		nums := []int{100, 1, 2, 100}
		expected := 200
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("example 1: [1,2,3,1]", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		expected := 4
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("example 2: [2,7,9,3,1]", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		expected := 12
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("single house", func(t *testing.T) {
		nums := []int{5}
		expected := 5
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("two houses", func(t *testing.T) {
		nums := []int{1, 2}
		expected := 2
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("all same values", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5}
		expected := 15 // rob houses at index 0, 2, 4
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("increasing values", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 9 // rob houses at index 0, 2, 4 = 1+3+5
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("decreasing values", func(t *testing.T) {
		nums := []int{5, 4, 3, 2, 1}
		expected := 9 // rob houses at index 0, 2, 4 = 5+3+1
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("large gap in middle", func(t *testing.T) {
		nums := []int{2, 1, 1, 2}
		expected := 4 // rob houses at index 0, 3 = 2+2
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("zeros in array", func(t *testing.T) {
		nums := []int{0, 0, 0}
		expected := 0
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("mixed with zeros", func(t *testing.T) {
		nums := []int{5, 0, 0, 5}
		expected := 10 // rob houses at index 0, 3
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("large values", func(t *testing.T) {
		nums := []int{100, 1, 1, 100}
		expected := 200 // rob houses at index 0, 3
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("three houses alternating", func(t *testing.T) {
		nums := []int{2, 1, 1}
		expected := 3 // rob houses at index 0, 2 = 2+1
		result := rob(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
