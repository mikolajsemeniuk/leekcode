package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/climbing-stairs
//
// 70. Climbing Stairs
//
// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
//
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
//
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// Constraints:
//
// 1 <= n <= 45
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// if n == 3 { return 4 }

	waysTwoSteps := 1 // n=1
	waysOneStep := 2  // n=2
	// waysThreeStepBack := 1 // n=1
	// waysTwoStepBack   := 2 // n=2
	// waysOneStepBack    := 4 // n=3

	for i := 3; i <= n; i++ {
		// current := waysOneStepBack + waysTwoStepBack + waysThreeStepBack
		current := waysOneStep + waysTwoSteps

		waysTwoSteps = waysOneStep
		waysOneStep = current
		// waysThreeStepBack = waysTwoStepBack
		// waysTwoStepBack = waysOneStepBack
		// waysOneStepBack = current
	}

	return waysOneStep
}

func TestClimbStairs(t *testing.T) {
	t.Run("n = 3", func(t *testing.T) {
		n := 3
		expected := 3
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 2", func(t *testing.T) {
		n := 2
		expected := 2
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 1", func(t *testing.T) {
		n := 1
		expected := 1
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 4", func(t *testing.T) {
		n := 4
		expected := 5
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 5", func(t *testing.T) {
		n := 5
		expected := 8
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 10", func(t *testing.T) {
		n := 10
		expected := 89
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 20", func(t *testing.T) {
		n := 20
		expected := 10946
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 30", func(t *testing.T) {
		n := 30
		expected := 1346269
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("n = 45", func(t *testing.T) {
		n := 45
		expected := 1836311903
		result := climbStairs(n)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
