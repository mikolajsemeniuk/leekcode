package leekcode

import (
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/subsets
//
// 78. Subsets
//
// Given an integer array nums of unique elements, return all possible subsets (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Example 1:
//
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
// Example 2:
//
// Input: nums = [0]
// Output: [[],[0]]
//
// Constraints:
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
func subsets(nums []int) [][]int {
	out := make([][]int, 0, 1<<len(nums))
	tmp := make([]int, 0, len(nums))

	var walk func(index int)
	walk = func(index int) {
		if index == len(nums) {
			in := make([]int, len(tmp))
			copy(in, tmp)
			out = append(out, in)
			return
		}

		walk(index + 1)
		tmp = append(tmp, nums[index])

		walk(index + 1)
		tmp = tmp[:len(tmp)-1]
	}

	walk(0)

	return out
}

// 1
// ├─ N → 2
// │   ├─ N → 3
// │   │   ├─ N → [] (empty)
// │   │   └─ B → [3]
// │   └─ B → 3
// │       ├─ N → [2]
// │       └─ B → [2,3]
// └─ B → 2
//
//	├─ N → 3
//	│   ├─ N → [1]
//	│   └─ B → [1,3]
//	└─ B → 3
//	    ├─ N → [1,2]
//	    └─ B → [1,2,3]
func TestSubsets(t *testing.T) {
	t.Run("nums = [1,2,3]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
		result := subsets(nums)
		if !equalSubsetSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [0]", func(t *testing.T) {
		nums := []int{0}
		expected := [][]int{{}, {0}}
		result := subsets(nums)
		if !equalSubsetSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nums = [4,5]", func(t *testing.T) {
		nums := []int{4, 5}
		expected := [][]int{{}, {4}, {5}, {4, 5}}
		result := subsets(nums)
		if !equalSubsetSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func equalSubsetSets(a, b [][]int) bool {
	normalize := func(sets [][]int) [][]int {
		out := make([][]int, 0, len(sets))
		for _, set := range sets {
			copySet := make([]int, len(set))
			copy(copySet, set)
			sort.Ints(copySet)
			out = append(out, copySet)
		}
		sort.Slice(out, func(i, j int) bool {
			if len(out[i]) != len(out[j]) {
				return len(out[i]) < len(out[j])
			}
			for k := 0; k < len(out[i]); k++ {
				if out[i][k] != out[j][k] {
					return out[i][k] < out[j][k]
				}
			}
			return false
		})
		return out
	}

	na := normalize(a)
	nb := normalize(b)
	return reflect.DeepEqual(na, nb)
}
