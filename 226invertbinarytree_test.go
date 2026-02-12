package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/invert-binary-tree
//
// 226. Invert Binary Tree
//
// Given the root of a binary tree, invert the tree, and return its root.
//
// Example 1:
//
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
//
// Example 2:
//
// Input: root = [2,1,3]
// Output: [2,3,1]
//
// Example 3:
//
// Input: root = []
// Output: []
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Helper function to convert a binary tree to level-order slice
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestInvertTree(t *testing.T) {
	t.Run("root = [4,2,7,1,3,6,9]", func(t *testing.T) {
		root := createTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
		expected := []interface{}{4, 7, 2, 9, 6, 3, 1}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [2,1,3]", func(t *testing.T) {
		root := createTree([]interface{}{2, 1, 3})
		expected := []interface{}{2, 3, 1}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = []", func(t *testing.T) {
		root := createTree([]interface{}{})
		expected := []interface{}{}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [1]", func(t *testing.T) {
		root := createTree([]interface{}{1})
		expected := []interface{}{1}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [1,2]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2})
		expected := []interface{}{1, nil, 2}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [1,nil,2]", func(t *testing.T) {
		root := createTree([]interface{}{1, nil, 2})
		expected := []interface{}{1, 2}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [1,2,3,4,5,6,7]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, 5, 6, 7})
		expected := []interface{}{1, 3, 2, 7, 6, 5, 4}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})

	t.Run("root = [1,2,3,nil,nil,4,5]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, nil, nil, 4, 5})
		expected := []interface{}{1, 3, 2, 5, 4}
		result := invertTree(root)
		resultSlice := treeToSlice(result)
		if !reflect.DeepEqual(resultSlice, expected) {
			t.Errorf("Expected %v, got %v", expected, resultSlice)
		}
	})
}
