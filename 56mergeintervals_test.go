package leekcode

import (
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/merge-intervals
//
// 56. Merge Intervals
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Example 1:
//
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
//
// Example 2:
//
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
//
// Constraints:
//
// 1 <= intervals.length <= 10^4
// intervals[i].length == 2
// 0 <= starti <= endi <= 10^4
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	out := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := out[len(out)-1]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			out = append(out, curr)
		}
	}

	return out
}

func TestMerge(t *testing.T) {
	t.Run("intervals = [[1,3],[2,6],[8,10],[15,18]]", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
		result := merge(intervals)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("intervals = [[1,4],[4,5]]", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {4, 5}}
		expected := [][]int{{1, 5}}
		result := merge(intervals)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("intervals = [[1,4],[0,2],[3,5]]", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {0, 2}, {3, 5}}
		expected := [][]int{{0, 5}}
		result := merge(intervals)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("intervals = [[1,4]]", func(t *testing.T) {
		intervals := [][]int{{1, 4}}
		expected := [][]int{{1, 4}}
		result := merge(intervals)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
