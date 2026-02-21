package leekcode

import (
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/3sum
//
// 15. Three Sum
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
// Example 1:
//
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
//
// Example 2:
//
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
//
// Example 3:
//
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
//
// Constraints:
//
// 3 <= nums.length <= 3000
// -10^5 <= nums[i] <= 10^5
func threeSum(nums []int) [][]int {
	var out [][]int

	if len(nums) < 3 {
		return out
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early termination: if the smallest number is positive, no solution exists
		if nums[i] > 0 {
			return out
		}

		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum < target {
				left++
				continue
			}

			if sum > target {
				right--
				continue
			}

			in := []int{nums[i], nums[left], nums[right]}
			out = append(out, in)

			// Skip duplicates for the second element
			for left < right && nums[left] == nums[left+1] {
				left++
			}

			// Skip duplicates for the third element
			for left < right && nums[right] == nums[right-1] {
				right--
			}

			left++
			right--
		}
	}

	return out
}

// Helper function to compare two sets of triplets (order-independent)
func sameTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each triplet and the overall list for comparison
	sortTriplets := func(triplets [][]int) [][]int {
		sorted := make([][]int, len(triplets))
		for i, triplet := range triplets {
			sorted[i] = make([]int, len(triplet))
			copy(sorted[i], triplet)
			sort.Ints(sorted[i])
		}
		sort.Slice(sorted, func(i, j int) bool {
			for k := 0; k < 3; k++ {
				if sorted[i][k] != sorted[j][k] {
					return sorted[i][k] < sorted[j][k]
				}
			}
			return false
		})
		return sorted
	}

	sortedA := sortTriplets(a)
	sortedB := sortTriplets(b)

	return reflect.DeepEqual(sortedA, sortedB)
}

func TestThreeSum(t *testing.T) {
	t.Run("example 1: [-1,0,1,2,-1,-4]", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		result := threeSum(nums)
		expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 2: [0,1,1]", func(t *testing.T) {
		nums := []int{0, 1, 1}
		result := threeSum(nums)
		expected := [][]int{}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 3: [0,0,0]", func(t *testing.T) {
		nums := []int{0, 0, 0}
		result := threeSum(nums)
		expected := [][]int{{0, 0, 0}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("no solution", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := threeSum(nums)
		expected := [][]int{}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("all negative", func(t *testing.T) {
		nums := []int{-1, -2, -3, -4, -5}
		result := threeSum(nums)
		expected := [][]int{}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single triplet", func(t *testing.T) {
		nums := []int{-1, 0, 1}
		result := threeSum(nums)
		expected := [][]int{{-1, 0, 1}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("multiple zeros", func(t *testing.T) {
		nums := []int{0, 0, 0, 0}
		result := threeSum(nums)
		expected := [][]int{{0, 0, 0}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("duplicates handled correctly", func(t *testing.T) {
		nums := []int{-1, -1, -1, 0, 1, 1, 1}
		result := threeSum(nums)
		expected := [][]int{{-1, 0, 1}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("multiple solutions", func(t *testing.T) {
		nums := []int{-2, 0, 0, 2, 2}
		result := threeSum(nums)
		expected := [][]int{{-2, 0, 2}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("large positive and negative", func(t *testing.T) {
		nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
		result := threeSum(nums)
		expected := [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("with zeros in middle", func(t *testing.T) {
		nums := []int{-1, 0, 1, 0}
		result := threeSum(nums)
		expected := [][]int{{-1, 0, 1}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("minimum length array", func(t *testing.T) {
		nums := []int{1, 2, -3}
		result := threeSum(nums)
		expected := [][]int{{-3, 1, 2}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("same number multiple times different combos", func(t *testing.T) {
		nums := []int{-2, -2, -2, -1, -1, -1, 0, 0, 0, 1, 1, 1, 2, 2, 2}
		result := threeSum(nums)
		expected := [][]int{{-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("edge case with two zeros", func(t *testing.T) {
		nums := []int{0, 0, 1}
		result := threeSum(nums)
		expected := [][]int{}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mixed values small array", func(t *testing.T) {
		nums := []int{3, 0, -2, -1, 1, 2}
		result := threeSum(nums)
		expected := [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}

		if !sameTriplets(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
