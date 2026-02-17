package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/product-of-array-except-self
//
// 238. Product of Array Except Self
//
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
//
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
//
// Example 2:
//
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
//
// Constraints:
//
// 2 <= nums.length <= 10^5
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// First pass: calculate prefix products
	// result[i] will contain the product of all elements to the left of i
	result[0] = 1
	for i := 1; i < n; i++ {
		x := result[i-1]
		y := nums[i-1]
		result[i] = x * y
	}

	// Second pass: calculate suffix products and multiply with prefix
	// Use a variable to track the product of all elements to the right
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		x := result[i]
		result[i] = x * suffixProduct
		y := nums[i]
		suffixProduct *= y
	}

	return result
}

func TestProductExceptSelf(t *testing.T) {
	t.Run("example 1: [1,2,3,4]", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 2: [-1,1,0,-3,3]", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("two elements", func(t *testing.T) {
		nums := []int{2, 3}
		expected := []int{3, 2}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("all ones", func(t *testing.T) {
		nums := []int{1, 1, 1, 1}
		expected := []int{1, 1, 1, 1}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("with negative numbers", func(t *testing.T) {
		nums := []int{-1, -2, -3, -4}
		expected := []int{-24, -12, -8, -6}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zero at start", func(t *testing.T) {
		nums := []int{0, 1, 2, 3}
		expected := []int{6, 0, 0, 0}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zero at end", func(t *testing.T) {
		nums := []int{1, 2, 3, 0}
		expected := []int{0, 0, 0, 6}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zero in middle", func(t *testing.T) {
		nums := []int{1, 2, 0, 4}
		expected := []int{0, 0, 8, 0}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("multiple zeros", func(t *testing.T) {
		nums := []int{0, 0, 1, 2}
		expected := []int{0, 0, 0, 0}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mixed positive and negative", func(t *testing.T) {
		nums := []int{2, -3, 4, -5}
		expected := []int{60, -40, 30, -24}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("large array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{3628800, 1814400, 1209600, 907200, 725760, 604800, 518400, 453600, 403200, 362880}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("powers of two", func(t *testing.T) {
		nums := []int{1, 2, 4, 8}
		expected := []int{64, 32, 16, 8}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single zero with others", func(t *testing.T) {
		nums := []int{5, 0, 3}
		expected := []int{0, 15, 0}
		result := productExceptSelf(nums)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
