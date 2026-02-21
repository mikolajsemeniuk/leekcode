package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/longest-increasing-subsequence
//
// 300. Longest Increasing Subsequence
//
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
// Example 1:
//
// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//
// Example 2:
//
// Input: nums = [0,1,0,3,2,3]
// Output: 4
//
// Example 3:
//
// Input: nums = [7,7,7,7,7,7,7]
// Output: 1
//
// Constraints:
//
// 1 <= nums.length <= 2500
// -10^4 <= nums[i] <= 10^4
//
// Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, num := range nums {
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}

func TestLengthOfLIS(t *testing.T) {
	t.Run("nums = [10,9,2,5,3,7,101,18]", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		expected := 4
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,1,0,3,2,3]", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 2, 3}
		expected := 4
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [7,7,7,7,7,7,7]", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7, 7}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [4,10,4,3,8,9]", func(t *testing.T) {
		nums := []int{4, 10, 4, 3, 8, 9}
		expected := 3
		result := lengthOfLIS(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
