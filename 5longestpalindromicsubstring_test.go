package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/longest-palindromic-substring
//
// 5. Longest Palindromic Substring
//
// Given a string s, return the longest palindromic substring in s.
//
// Example 1:
//
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
//
// Example 2:
//
// Input: s = "cbbd"
// Output: "bb"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters.
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1
			if length > maxLen {
				maxLen = length
				start = left
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return s[start : start+maxLen]
}

func TestLongestPalindrome(t *testing.T) {
	t.Run(`s = "babad"`, func(t *testing.T) {
		s := "babad"
		result := longestPalindrome(s)
		if result != "bab" && result != "aba" {
			t.Errorf("Expected %q or %q, got %q", "bab", "aba", result)
		}
	})

	t.Run(`s = "cbbd"`, func(t *testing.T) {
		s := "cbbd"
		expected := "bb"
		result := longestPalindrome(s)
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run(`s = "a"`, func(t *testing.T) {
		s := "a"
		expected := "a"
		result := longestPalindrome(s)
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run(`s = "ac"`, func(t *testing.T) {
		s := "ac"
		result := longestPalindrome(s)
		if result != "a" && result != "c" {
			t.Errorf("Expected %q or %q, got %q", "a", "c", result)
		}
	})

	t.Run(`s = "forgeeksskeegfor"`, func(t *testing.T) {
		s := "forgeeksskeegfor"
		expected := "geeksskeeg"
		result := longestPalindrome(s)
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})
}
