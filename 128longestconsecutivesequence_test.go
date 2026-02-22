package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/longest-consecutive-sequence
//
// 128. Longest Consecutive Sequence
//
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
// Example 1:
//
// Input: nums = [100,4,200,1,3,2]
// Output: 4
//
// Example 2:
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
//
// Constraints:
//
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for n := range set {
		if _, exists := set[n-1]; exists {
			continue
		}

		length := 1
		for current := n + 1; ; current++ {
			if _, exists := set[current]; !exists {
				break
			}
			length++
		}

		longest = max(longest, length)
	}

	return longest
}

func TestLongestConsecutive(t *testing.T) {
	t.Run("nums = [100,4,200,1,3,2]", func(t *testing.T) {
		nums := []int{100, 4, 200, 1, 3, 2}
		expected := 4
		result := longestConsecutive(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [0,3,7,2,5,8,4,6,0,1]", func(t *testing.T) {
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		expected := 9
		result := longestConsecutive(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = []", func(t *testing.T) {
		nums := []int{}
		expected := 0
		result := longestConsecutive(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [1]", func(t *testing.T) {
		nums := []int{1}
		expected := 1
		result := longestConsecutive(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("nums = [9,1,4,7,3,-1,0,5,8,-1,6]", func(t *testing.T) {
		nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
		expected := 7
		result := longestConsecutive(nums)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
