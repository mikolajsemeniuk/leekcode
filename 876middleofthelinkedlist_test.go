package leekcode

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/middle-of-the-linked-list/
//
// 876. Middle of the Linked List
//
// Given the head of a singly linked list, return the middle node of the linked list.
//
// If there are two middle nodes, return the second middle node.
//
// Example 1:
//
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.
// Example 2:
//
// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
//
// Constraints:
//
// The number of nodes in the list is in the range [1, 100].
// 1 <= Node.val <= 100
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func TestMiddleNode(t *testing.T) {
	t.Run("head = [1,2,3,4,5]", func(t *testing.T) {
		head := createList([]int{1, 2, 3, 4, 5})
		expected := []int{3, 4, 5}
		result := middleNode(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2,3,4,5,6]", func(t *testing.T) {
		head := createList([]int{1, 2, 3, 4, 5, 6})
		expected := []int{4, 5, 6}
		result := middleNode(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1]", func(t *testing.T) {
		head := createList([]int{1})
		expected := []int{1}
		result := middleNode(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2]", func(t *testing.T) {
		head := createList([]int{1, 2})
		expected := []int{2}
		result := middleNode(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("head = [1,2,3]", func(t *testing.T) {
		head := createList([]int{1, 2, 3})
		expected := []int{2, 3}
		result := middleNode(head)
		resultSlice := listToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})
}
