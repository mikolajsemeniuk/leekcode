package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/path-sum
//
// 112. Path Sum
//
// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path
// such that adding up all the values along the path equals targetSum.
//
// A leaf is a node with no children.
//
// Example 1:
//
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.
//
// Example 2:
//
// Input: root = [1,2,3], targetSum = 5
// Output: false
// Explanation: There two root-to-leaf paths in the tree:
// (1 --> 2): The sum is 3.
// (1 --> 3): The sum is 4.
// There is no root-to-leaf path with sum = 5.
//
// Example 3:
//
// Input: root = [], targetSum = 0
// Output: false
// Explanation: Since the tree is empty, there are no root-to-leaf paths.
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	target := targetSum - root.Val
	left := hasPathSum(root.Left, target)
	right := hasPathSum(root.Right, target)
	return left || right
}

func TestHasPathSum(t *testing.T) {
	t.Run("root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22", func(t *testing.T) {
		root := createTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1})
		targetSum := 22
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3], targetSum = 5", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3})
		targetSum := 5
		expected := false
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [], targetSum = 0", func(t *testing.T) {
		root := createTree([]interface{}{})
		targetSum := 0
		expected := false
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2], targetSum = 1", func(t *testing.T) {
		root := createTree([]interface{}{1, 2})
		targetSum := 1
		expected := false
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1], targetSum = 1", func(t *testing.T) {
		root := createTree([]interface{}{1})
		targetSum := 1
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2], targetSum = 3", func(t *testing.T) {
		root := createTree([]interface{}{1, 2})
		targetSum := 3
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [-2,null,-3], targetSum = -5", func(t *testing.T) {
		root := createTree([]interface{}{-2, nil, -3})
		targetSum := -5
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3,4,5], targetSum = 7", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, 5})
		targetSum := 7
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,-2,-3,1,3,-2,null,-1], targetSum = -1", func(t *testing.T) {
		root := createTree([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
		targetSum := -1
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,-2,-3,1,3,-2,null,-1], targetSum = 2", func(t *testing.T) {
		root := createTree([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
		targetSum := 2
		expected := true
		result := hasPathSum(root, targetSum)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
