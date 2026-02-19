package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal
//
// 102. Binary Tree Level Order Traversal
//
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
//
// Example 1:
//
// Input: root = [3,9,20,nil,nil,15,7]
// Output: [[3],[9,20],[15,7]]
//
// Example 2:
//
// Input: root = [1]
// Output: [[1]]
//
// Example 3:
//
// Input: root = []
// Output: []
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000
func levelOrder(root *TreeNode) [][]int {
	var out [][]int

	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}

		if depth == len(out) {
			out = append(out, []int{})
		}

		out[depth] = append(out[depth], n.Val)

		dfs(n.Left, depth+1)
		dfs(n.Right, depth+1)
	}

	dfs(root, 0)
	return out
}

func TestLevelOrder(t *testing.T) {
	t.Run("root = [3,9,20,nil,nil,15,7]", func(t *testing.T) {
		root := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
		expected := [][]int{{3}, {9, 20}, {15, 7}}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1]", func(t *testing.T) {
		root := createTree([]interface{}{1})
		expected := [][]int{{1}}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = []", func(t *testing.T) {
		root := createTree([]interface{}{})
		expected := [][]int{}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("root = [1,2,3,4,nil,nil,5]", func(t *testing.T) {
		root := createTree([]interface{}{1, 2, 3, 4, nil, nil, 5})
		expected := [][]int{{1}, {2, 3}, {4, 5}}
		result := levelOrder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
