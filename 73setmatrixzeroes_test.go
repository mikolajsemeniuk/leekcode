package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/set-matrix-zeroes
//
// 73. Set Matrix Zeroes
//
// Given an m x n integer matrix, if an element is 0, set its entire row and column to 0's.
//
// You must do it in place.
//
// Example 1:
//
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
//
// Example 2:
//
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// Constraints:
//
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2^31 <= matrix[i][j] <= 2^31 - 1
//
// Follow up:
// A straightforward solution using O(m*n) space is probably a bad idea.
// A simple improvement uses O(m+n) space, but still not the best solution.
// Could you devise a constant space solution?
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	zeroRows := make([]bool, rows)
	zeroCols := make([]bool, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				zeroRows[r] = true
				zeroCols[c] = true
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if zeroRows[r] || zeroCols[c] {
				matrix[r][c] = 0
			}
		}
	}
}

func TestSetZeroes(t *testing.T) {
	t.Run("matrix = [[1,1,1],[1,0,1],[1,1,1]]", func(t *testing.T) {
		matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
		setZeroes(matrix)
		if !reflect.DeepEqual(matrix, expected) {
			t.Errorf("Expected %v, got %v", expected, matrix)
		}
	})

	t.Run("matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]", func(t *testing.T) {
		matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
		expected := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
		setZeroes(matrix)
		if !reflect.DeepEqual(matrix, expected) {
			t.Errorf("Expected %v, got %v", expected, matrix)
		}
	})

	t.Run("matrix = [[1,2,3,4]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3, 4}}
		expected := [][]int{{1, 2, 3, 4}}
		setZeroes(matrix)
		if !reflect.DeepEqual(matrix, expected) {
			t.Errorf("Expected %v, got %v", expected, matrix)
		}
	})

	t.Run("matrix = [[1],[0],[3]]", func(t *testing.T) {
		matrix := [][]int{{1}, {0}, {3}}
		expected := [][]int{{0}, {0}, {0}}
		setZeroes(matrix)
		if !reflect.DeepEqual(matrix, expected) {
			t.Errorf("Expected %v, got %v", expected, matrix)
		}
	})
}
