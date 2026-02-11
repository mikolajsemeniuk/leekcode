package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/reverse-linked-list
//
// 206. Reverse Linked List
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Example 1:
//
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Example 2:
//
// Input: head = [1,2]
// Output: [2,1]
//
// Example 3:
//
// Input: head = []
// Output: []
//
// Constraints:
//
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000
//
// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func TestReverseList(t *testing.T) {
	t.Run("head = [1,2,3,4,5]", func(t *testing.T) {
		head := createList([]int{1, 2, 3, 4, 5})
		expected := []int{5, 4, 3, 2, 1}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2]", func(t *testing.T) {
		head := createList([]int{1, 2})
		expected := []int{2, 1}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = []", func(t *testing.T) {
		head := createList([]int{})
		expected := []int{}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1]", func(t *testing.T) {
		head := createList([]int{1})
		expected := []int{1}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2,3]", func(t *testing.T) {
		head := createList([]int{1, 2, 3})
		expected := []int{3, 2, 1}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2,3,4,5,6,7,8,9,10]", func(t *testing.T) {
		head := createList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		result := reverseList(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})
}
