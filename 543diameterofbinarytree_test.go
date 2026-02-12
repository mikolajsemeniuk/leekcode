package leekcode

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/diameter-of-binary-tree
//
// 543. Diameter of Binary Tree
//
// Given the root of a binary tree, return the length of the diameter of the tree.
//
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges between them.
//
// Example 1:
//
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
//
// Example 2:
//
// Input: root = [1,2]
// Output: 1
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 10^4].
// -100 <= Node.val <= 100
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		current := left + right
		if current > diameter {
			diameter = current
		}

		return max(left, right) + 1
	}

	depth(root)
	return diameter
}

// Helper function to create a binary tree from a level-order array
// nil values in the array represent missing nodes
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("root = [1,2,3,4,5]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, 5})
		expected := 3
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2})
		expected := 1
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1]", func(t *testing.T) {
		root := createTree([]interface{}{1})
		expected := 0
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3,4,5,nil,nil]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, 5, nil, nil})
		expected := 3
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3,nil,nil,4,5]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, nil, nil, 4, 5})
		expected := 3
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,nil,3,nil,4]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, nil, 3, nil, 4})
		expected := 3
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3,4,5,6,7]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, 5, 6, 7})
		expected := 4
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [4,-7,-3,nil,nil,-9,-3,9,-7,-4,nil,6,nil,-6,-6,nil,nil,0,6,5,nil,9,nil,nil,-1,-4,nil,nil,nil,-2]", func(t *testing.T) {
		root := createTree([]interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2})
		expected := 8
		result := diameterOfBinaryTree(root)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
