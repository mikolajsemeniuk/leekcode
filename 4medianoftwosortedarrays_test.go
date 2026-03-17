package leekcode

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays
//
// 4. Median of Two Sorted Arrays
//
// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
// Example 1:
//
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
//
// Example 2:
//
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
//
// Constraints:
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}

func findMedianSortedArraysSlower(nums1 []int, nums2 []int) float64 {
	in := append(nums1, nums2...)
	sort.Ints(in)

	n := len(in)
	if n%2 != 0 {
		return float64(in[n/2])
	}

	mid1 := in[n/2-1]
	mid2 := in[n/2]

	return float64(mid1+mid2) / 2.0
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("nums1 = [1,3], nums2 = [2]", func(t *testing.T) {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		expected := 2.0
		result := findMedianSortedArrays(nums1, nums2)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("nums1 = [1,2], nums2 = [3,4]", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		expected := 2.5
		result := findMedianSortedArrays(nums1, nums2)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("nums1 = [], nums2 = [1]", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{1}
		expected := 1.0
		result := findMedianSortedArrays(nums1, nums2)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("nums1 = [2], nums2 = []", func(t *testing.T) {
		nums1 := []int{2}
		nums2 := []int{}
		expected := 2.0
		result := findMedianSortedArrays(nums1, nums2)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("nums1 = [1,2,3], nums2 = [4,5,6,7,8]", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{4, 5, 6, 7, 8}
		expected := 4.5
		result := findMedianSortedArrays(nums1, nums2)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})
}
