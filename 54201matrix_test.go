package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/01-matrix
//
// 542. 01 Matrix
//
// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
//
// The distance between two adjacent cells is 1.
//
// Example 1:
//
// Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
// Output: [[0,0,0],[0,1,0],[0,0,0]]
//
// Example 2:
//
// Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
// Output: [[0,0,0],[0,1,0],[1,2,1]]
//
// Constraints:
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10^4
// 1 <= m * n <= 10^4
// mat[i][j] is either 0 or 1.
// There is at least one 0 in mat.
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	rows, cols := len(mat), len(mat[0])
	dist := make([][]int, rows)
	queue := make([][2]int, 0, rows*cols)

	for r := 0; r < rows; r++ {
		dist[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				dist[r][c] = 0
				queue = append(queue, [2]int{r, c})
			} else {
				dist[r][c] = -1
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	head := 0
	for head < len(queue) {
		cell := queue[head]
		head++
		r, c := cell[0], cell[1]

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if dist[nr][nc] != -1 {
				continue
			}
			dist[nr][nc] = dist[r][c] + 1
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return dist
}

func TestUpdateMatrix(t *testing.T) {
	t.Run("mat = [[0,0,0],[0,1,0],[0,0,0]]", func(t *testing.T) {
		mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
		expected := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
		result := updateMatrix(mat)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mat = [[0,0,0],[0,1,0],[1,1,1]]", func(t *testing.T) {
		mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
		expected := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}
		result := updateMatrix(mat)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mat = [[1,0,1],[1,1,1],[1,1,0]]", func(t *testing.T) {
		mat := [][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 0}}
		expected := [][]int{{1, 0, 1}, {2, 1, 1}, {2, 1, 0}}
		result := updateMatrix(mat)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mat = [[0]]", func(t *testing.T) {
		mat := [][]int{{0}}
		expected := [][]int{{0}}
		result := updateMatrix(mat)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("mat = [[0,1,1],[1,1,1]]", func(t *testing.T) {
		mat := [][]int{{0, 1, 1}, {1, 1, 1}}
		expected := [][]int{{0, 1, 2}, {1, 2, 3}}
		result := updateMatrix(mat)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
