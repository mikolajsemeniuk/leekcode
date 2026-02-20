package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/add-two-numbers
//
// 2. Add Two Numbers
//
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example 1:
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
//
// Example 2:
//
// Input: l1 = [0], l2 = [0]
// Output: [0]
//
// Example 3:
//
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
//
// Constraints:
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy head to simplify edge cases
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// Helper function to compare two slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAddTwoNumbers(t *testing.T) {
	t.Run("example 1: [2,4,3] + [5,6,4]", func(t *testing.T) {
		l1 := createList([]int{2, 4, 3})
		l2 := createList([]int{5, 6, 4})
		expected := []int{7, 0, 8}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("example 2: [0] + [0]", func(t *testing.T) {
		l1 := createList([]int{0})
		l2 := createList([]int{0})
		expected := []int{0}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("example 3: [9,9,9,9,9,9,9] + [9,9,9,9]", func(t *testing.T) {
		l1 := createList([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := createList([]int{9, 9, 9, 9})
		expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("single digit no carry", func(t *testing.T) {
		l1 := createList([]int{2})
		l2 := createList([]int{3})
		expected := []int{5}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("single digit with carry", func(t *testing.T) {
		l1 := createList([]int{5})
		l2 := createList([]int{5})
		expected := []int{0, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("different lengths l1 longer", func(t *testing.T) {
		l1 := createList([]int{1, 2, 3, 4})
		l2 := createList([]int{5, 6})
		expected := []int{6, 8, 3, 4}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("different lengths l2 longer", func(t *testing.T) {
		l1 := createList([]int{5, 6})
		l2 := createList([]int{1, 2, 3, 4})
		expected := []int{6, 8, 3, 4}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("carry propagates through list", func(t *testing.T) {
		l1 := createList([]int{9, 9, 9})
		l2 := createList([]int{1})
		expected := []int{0, 0, 0, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("carry at end", func(t *testing.T) {
		l1 := createList([]int{9, 9})
		l2 := createList([]int{1})
		expected := []int{0, 0, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("all zeros", func(t *testing.T) {
		l1 := createList([]int{0, 0, 0})
		l2 := createList([]int{0, 0, 0})
		expected := []int{0, 0, 0}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("multiple carries", func(t *testing.T) {
		l1 := createList([]int{9, 8, 7})
		l2 := createList([]int{1, 2, 3})
		expected := []int{0, 1, 1, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("two digit numbers", func(t *testing.T) {
		l1 := createList([]int{9, 1}) // 19
		l2 := createList([]int{1, 1}) // 11
		expected := []int{0, 3}       // 30
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("one list much longer", func(t *testing.T) {
		l1 := createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		l2 := createList([]int{5, 6, 4})
		expected := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("max single digits", func(t *testing.T) {
		l1 := createList([]int{9})
		l2 := createList([]int{9})
		expected := []int{8, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("alternating carries", func(t *testing.T) {
		l1 := createList([]int{5, 5})
		l2 := createList([]int{5, 5})
		expected := []int{0, 1, 1}
		result := addTwoNumbers(l1, l2)
		actual := listToSlice(result)
		if !slicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}
