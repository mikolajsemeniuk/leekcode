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
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			result = append(result, matrix[top][c])
		}
		top++

		for r := top; r <= bottom; r++ {
			result = append(result, matrix[r][right])
		}
		right--

		if top <= bottom {
			for c := right; c >= left; c-- {
				result = append(result, matrix[bottom][c])
			}
			bottom--
		}

		if left <= right {
			for r := bottom; r >= top; r-- {
				result = append(result, matrix[r][left])
			}
			left++
		}
	}

	return result
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
