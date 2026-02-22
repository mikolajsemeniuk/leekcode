package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/decode-ways
//
// 91. Decode Ways
//
// A message containing letters from A-Z can be encoded into numbers using the following mapping:
//
// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
//
// Given a string s containing only digits, return the number of ways to decode it.
//
// Example 1:
//
// Input: s = "12"
// Output: 2
// Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
//
// Example 2:
//
// Input: s = "226"
// Output: 3
// Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
//
// Example 3:
//
// Input: s = "06"
// Output: 0
// Explanation: "06" cannot be mapped to "F" because of the leading zero.
//
// Constraints:
//
// 1 <= s.length <= 100
// s contains only digits and may contain leading zeros.
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev2 := 1
	prev1 := 1

	for i := 1; i < len(s); i++ {
		current := 0
		if s[i] != '0' {
			current += prev1
		}

		twoDigit := (s[i-1]-'0')*10 + (s[i] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			current += prev2
		}

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func TestNumDecodings(t *testing.T) {
	t.Run(`s = "12"`, func(t *testing.T) {
		s := "12"
		expected := 2
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run(`s = "226"`, func(t *testing.T) {
		s := "226"
		expected := 3
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run(`s = "06"`, func(t *testing.T) {
		s := "06"
		expected := 0
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run(`s = "10"`, func(t *testing.T) {
		s := "10"
		expected := 1
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run(`s = "27"`, func(t *testing.T) {
		s := "27"
		expected := 1
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run(`s = "11106"`, func(t *testing.T) {
		s := "11106"
		expected := 2
		result := numDecodings(s)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
