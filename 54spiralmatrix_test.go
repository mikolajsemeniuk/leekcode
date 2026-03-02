package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/spiral-matrix
//
// 54. Spiral Matrix
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
//
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	top, bottom := 0, rows-1
	left, right := 0, cols-1

	var out []int
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			out = append(out, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			out = append(out, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				out = append(out, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				out = append(out, matrix[i][left])
			}
			left++
		}
	}

	return out
}

func TestSpiralOrder(t *testing.T) {
	t.Run("matrix = [[1,2,3],[4,5,6],[7,8,9]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("matrix = [[1]]", func(t *testing.T) {
		matrix := [][]int{{1}}
		expected := []int{1}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("matrix = [[1,2,3,4]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3, 4}}
		expected := []int{1, 2, 3, 4}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("matrix = [[1],[2],[3],[4]]", func(t *testing.T) {
		matrix := [][]int{{1}, {2}, {3}, {4}}
		expected := []int{1, 2, 3, 4}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
