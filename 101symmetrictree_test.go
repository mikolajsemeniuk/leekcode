package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/symmetric-tree
//
// 101. Symmetric Tree
//
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Example 1:
//
// Input: root = [1,2,2,3,4,4,3]
// Output: true
//
// Example 2:
//
// Input: root = [1,2,2,null,3,null,3]
// Output: false
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
//
// Follow up: Could you solve it both recursively and iteratively?
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func TestIsSymmetric(t *testing.T) {
	t.Run("root = [1,2,2,3,4,4,3]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,2,null,3,null,3]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
		expected := false
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1]", func(t *testing.T) {
		root := createTree([]interface{}{1})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,2]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3})
		expected := false
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,2,3,4,4,3,5,6,7,8,8,7,6,5]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,2,null,3,3,null]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2, nil, 3, 3, nil})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,2,2,null,2]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 2, 2, nil, 2})
		expected := false
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,0,0]", func(t *testing.T) {
		root := createTree([]interface{}{1, 0, 0})
		expected := true
		result := isSymmetric(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
