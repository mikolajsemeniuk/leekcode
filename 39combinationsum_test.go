package leekcode

import (
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/combination-sum
//
// 39. Combination Sum
//
// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers sum to target.
//
// You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times.
//
// Example 1:
//
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
//
// Example 2:
//
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]
//
// Example 3:
//
// Input: candidates = [2], target = 1
// Output: []
//
// Constraints:
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var out [][]int
	var tmp []int

	var walk func(start, left int)
	walk = func(start, left int) {
		if left == 0 {
			in := make([]int, len(tmp))
			copy(in, tmp)
			out = append(out, in)
			return
		}

		if left < 0 {
			return
		}

		// Start from the current index to prevent duplicate combinations by never looking backward.
		for i := start; i < len(candidates); i++ {
			val := candidates[i]
			if val > left {
				break
			}

			tmp = append(tmp, val)
			// Re-use the current index 'i' to allow the same candidate to be chosen multiple times.
			walk(i, left-val)
			tmp = tmp[:len(tmp)-1]
		}
	}

	walk(0, target)

	return out
}

// start (rem=7)
// ├─ 2  -> rem=5
// │  ├─ 2 -> rem=3
// │  │  ├─ 2 -> rem=1
// │  │  │  ├─ 2 -> rem=-1 (return)
// │  │  │  ├─ 3 -> rem=-2 (break)
// │  │  │  ├─ 6 -> rem=-5 (break)
// │  │  │  └─ 7 -> rem=-6 (break)
// │  │  ├─ 3 -> rem=0  ✅ [2,2,3]
// │  │  ├─ 6 -> rem=-3 (break)
// │  │  └─ 7 -> rem=-4 (break)
// │  ├─ 3 -> rem=2
// │  │  ├─ 3 -> rem=-1 (break)
// │  │  ├─ 6 -> rem=-4 (break)
// │  │  └─ 7 -> rem=-5 (break)
// │  ├─ 6 -> rem=-1 (break)
// │  └─ 7 -> rem=-2 (break)
// ├─ 3 -> rem=4
// │  ├─ 3 -> rem=1
// │  │  ├─ 3 -> rem=-2 (break)
// │  │  ├─ 6 -> rem=-5 (break)
// │  │  └─ 7 -> rem=-6 (break)
// │  ├─ 6 -> rem=-2 (break)
// │  └─ 7 -> rem=-3 (break)
// ├─ 6 -> rem=1
// │  ├─ 6 -> rem=-5 (break)
// │  └─ 7 -> rem=-6 (break)
// └─ 7 -> rem=0 ✅ [7]
func TestCombinationSum(t *testing.T) {
	t.Run("candidates = [2,3,6,7], target = 7", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		expected := [][]int{{2, 2, 3}, {7}}
		result := combinationSum(candidates, target)
		if !equalCombinationSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("candidates = [2,3,5], target = 8", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
		result := combinationSum(candidates, target)
		if !equalCombinationSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("candidates = [2], target = 1", func(t *testing.T) {
		candidates := []int{2}
		target := 1
		expected := [][]int{}
		result := combinationSum(candidates, target)
		if !equalCombinationSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("candidates = [3,4,7], target = 14", func(t *testing.T) {
		candidates := []int{3, 4, 7}
		target := 14
		expected := [][]int{{3, 4, 7}, {7, 7}, {3, 3, 4, 4}}
		result := combinationSum(candidates, target)
		if !equalCombinationSets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func equalCombinationSets(a, b [][]int) bool {
	normalize := func(combos [][]int) [][]int {
		out := make([][]int, 0, len(combos))
		for _, combo := range combos {
			c := make([]int, len(combo))
			copy(c, combo)
			sort.Ints(c)
			out = append(out, c)
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
