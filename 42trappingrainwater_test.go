package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/trapping-rain-water
//
// 42. Trapping Rain Water
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it can trap after raining.
//
// Example 1:
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
//
// Example 2:
//
// Input: height = [4,2,0,3,2,5]
// Output: 9
//
// Constraints:
//
// n == height.length
// 1 <= n <= 2 * 10^4
// 0 <= height[i] <= 10^5
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		// Wybieramy stronę, która jest niższa – to ona limituje wodę
		if height[left] < height[right] {
			// Jeśli obecny słupek jest wyższy/równy maxowi, aktualizujemy i skipujemy
			if height[left] >= leftMax {
				leftMax = height[left]
				left++
				continue
			}
			// Skoro nie weszło w continue, to znaczy, że możemy nalać wody
			water += leftMax - height[left]
			left++
		} else {
			// Analogicznie dla prawej strony
			if height[right] >= rightMax {
				rightMax = height[right]
				right--
				continue
			}
			water += rightMax - height[right]
			right--
		}
	}

	return water
}

func TestTrap(t *testing.T) {
	t.Run("height = [0,1,0,2,1,0,1,3,2,1,2,1]", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		expected := 6
		result := trap(height)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("height = [4,2,0,3,2,5]", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		expected := 9
		result := trap(height)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("height = [1,1,1]", func(t *testing.T) {
		height := []int{1, 1, 1}
		expected := 0
		result := trap(height)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("height = [2,0,2]", func(t *testing.T) {
		height := []int{2, 0, 2}
		expected := 2
		result := trap(height)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
