package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/coin-change
//
// 322. Coin Change
//
// You are given an integer array coins representing coins of different denominations
// and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
// Example 1:
//
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// Example 2:
//
// Input: coins = [2], amount = 3
// Output: -1
//
// Example 3:
//
// Input: coins = [1], amount = 0
// Output: 0
//
// Constraints:
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2^31 - 1
// 0 <= amount <= 10^4
// func coinChange(coins []int, amount int) int {
// 	if amount == 0 {
// 		return 0
// 	}

// 	slices.Sort(coins)
// 	slices.Reverse(coins)

// 	out := 0
// 	total := amount
// 	for _, v := range coins {
// 		x := total / v
// 		if x > 0 {
// 			total -= x * v
// 			out += x
// 		}

// 		if total == 0 {
// 			return out
// 		}
// 	}

// 	return -1
// }

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	max := amount + 1
	out := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		out[i] = max
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && out[i-coin]+1 < out[i] {
				out[i] = out[i-coin] + 1
			}
		}
	}

	if out[amount] > amount {
		return -1
	}

	return out[amount]
}

func TestCoinChange(t *testing.T) {
	t.Run("coins = [1,2,5], amount = 11", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 11
		expected := 3
		result := coinChange(coins, amount)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("coins = [2], amount = 3", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		expected := -1
		result := coinChange(coins, amount)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("coins = [1], amount = 0", func(t *testing.T) {
		coins := []int{1}
		amount := 0
		expected := 0
		result := coinChange(coins, amount)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("coins = [2,5,10,1], amount = 27", func(t *testing.T) {
		coins := []int{2, 5, 10, 1}
		amount := 27
		expected := 4
		result := coinChange(coins, amount)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("coins = [2,4], amount = 7", func(t *testing.T) {
		coins := []int{2, 4}
		amount := 7
		expected := -1
		result := coinChange(coins, amount)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
