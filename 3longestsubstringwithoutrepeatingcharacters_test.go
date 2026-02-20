package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters
//
// 3. Longest Substring Without Repeating Characters
//
// Given a string s, find the length of the longest substring without repeating characters.
//
// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
// Constraints:
//
// 0 <= s.length <= 5 * 10^4
// s consists of English letters, digits, symbols and spaces.
func lengthOfLongestSubstring(s string) int {
	start := 0
	total := 0
	m := map[rune]int{}

	for i, char := range s {
		if val, exists := m[char]; exists && val >= start {
			start = val + 1
		}

		m[char] = i

		current := i - start + 1
		total = max(total, current)
	}

	return total
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("example 1: abcabcbb", func(t *testing.T) {
		s := "abcabcbb"
		expected := 3
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("example 2: bbbbb", func(t *testing.T) {
		s := "bbbbb"
		expected := 1
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("example 3: pwwkew", func(t *testing.T) {
		s := "pwwkew"
		expected := 3
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		s := ""
		expected := 0
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("single character", func(t *testing.T) {
		s := "a"
		expected := 1
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("all unique characters", func(t *testing.T) {
		s := "abcdef"
		expected := 6
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("two characters alternating", func(t *testing.T) {
		s := "abababab"
		expected := 2
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("repeating at end", func(t *testing.T) {
		s := "abcdefgg"
		expected := 7
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("repeating at start", func(t *testing.T) {
		s := "aabcdef"
		expected := 6
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("with spaces", func(t *testing.T) {
		s := "a b c a"
		expected := 3 // "a b" or " bc"
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("with digits", func(t *testing.T) {
		s := "12345612"
		expected := 6
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("with symbols", func(t *testing.T) {
		s := "a!b@c#a"
		expected := 6
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("longest in middle", func(t *testing.T) {
		s := "abcdefghijklmnopqrstuvwxyz"
		expected := 26
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("complex pattern", func(t *testing.T) {
		s := "dvdf"
		expected := 3
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("repeated character appears later", func(t *testing.T) {
		s := "abcabcbb"
		expected := 3
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("special case: tmmzuxt", func(t *testing.T) {
		s := "tmmzuxt"
		expected := 5
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("consecutive duplicates then unique", func(t *testing.T) {
		s := "aabaab!bb"
		expected := 3
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("all spaces", func(t *testing.T) {
		s := "     "
		expected := 1
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
