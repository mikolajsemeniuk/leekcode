package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/linked-list-cycle
//
// 141. Linked List Cycle
//
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// Example 1:
//
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
//
// Example 2:
//
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
//
// Example 3:
//
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
//
// Constraints:
//
// The number of the nodes in the list is in the range [0, 10^4].
// -10^5 <= Node.val <= 10^5
// pos is -1 or a valid index in the linked-list.
//
// Follow up: Can you solve it using O(1) (i.e. constant) memory?
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func TestHasCycle(t *testing.T) {
	createListWithCycle := func(values []int, pos int) *ListNode {
		if len(values) == 0 {
			return nil
		}

		head := &ListNode{Val: values[0]}
		current := head
		var cycleNode *ListNode

		if pos == 0 {
			cycleNode = head
		}

		for i := 1; i < len(values); i++ {
			current.Next = &ListNode{Val: values[i]}
			current = current.Next

			if i == pos {
				cycleNode = current
			}
		}

		if pos >= 0 && cycleNode != nil {
			current.Next = cycleNode
		}

		return head
	}

	t.Run("head = [3,2,0,-4], pos = 1", func(t *testing.T) {
		head := createListWithCycle([]int{3, 2, 0, -4}, 1)
		expected := true
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [1,2], pos = 0", func(t *testing.T) {
		head := createListWithCycle([]int{1, 2}, 0)
		expected := true
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [1], pos = -1", func(t *testing.T) {
		head := createListWithCycle([]int{1}, -1)
		expected := false
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [], pos = -1", func(t *testing.T) {
		head := createListWithCycle([]int{}, -1)
		expected := false
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [1,2,3,4,5], pos = -1", func(t *testing.T) {
		head := createListWithCycle([]int{1, 2, 3, 4, 5}, -1)
		expected := false
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [1,2,3,4,5], pos = 2", func(t *testing.T) {
		head := createListWithCycle([]int{1, 2, 3, 4, 5}, 2)
		expected := true
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("head = [1,2], pos = -1", func(t *testing.T) {
		head := createListWithCycle([]int{1, 2}, -1)
		expected := false
		result := hasCycle(head)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
