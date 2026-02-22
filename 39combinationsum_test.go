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
	var result [][]int
	var current []int

	var dfs func(start, remaining int)
	dfs = func(start, remaining int) {
		if remaining == 0 {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			value := candidates[i]
			if value > remaining {
				break
			}
			current = append(current, value)
			dfs(i, remaining-value)
			current = current[:len(current)-1]
		}
	}

	dfs(0, target)
	return result
}

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
