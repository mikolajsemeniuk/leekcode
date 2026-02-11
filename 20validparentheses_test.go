package leekcode

import "testing"

// https://leetcode.com/problems/valid-parentheses/description/
//
// 20. Valid Parentheses
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
//
// Output: true
//
// Example 2:
//
// Input: s = "()[]{}"
//
// Output: true
//
// Example 3:
//
// Input: s = "(]"
//
// Output: false
//
// Example 4:
//
// Input: s = "([])"
//
// Output: true
//
// Example 5:
//
// Input: s = "([)]"
//
// Output: false
//
// Constraints:
//
// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.
func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		val, exists := pairs[char]
		if !exists {
			stack = append(stack, char)
			continue
		}

		index := len(stack) - 1
		if len(stack) == 0 || stack[index] != val {
			return false
		}

		stack = stack[:index]
	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	t.Run("()", func(t *testing.T) {
		if isValid("()") {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("()[]{}", func(t *testing.T) {
		if !isValid("()[]{}") {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("(]", func(t *testing.T) {
		if isValid("(]") {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("([)]", func(t *testing.T) {
		if isValid("([)]") {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("{[]}", func(t *testing.T) {
		if !isValid("{[]}") {
			t.Errorf("Expected true, got false")
		}
	})
}
