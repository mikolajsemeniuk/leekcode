package leekcode

import (
	"testing"
)

// https://leetcode.com/problems/number-of-islands
//
// 200. Number of Islands
//
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
//
// You may assume all four edges of the grid are all surrounded by water.
//
// Example 1:
//
// Input: grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// Output: 1
//
// Example 2:
//
// Input: grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// Output: 3
//
// Constraints:
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	var visit func(r, c int)
	visit = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}

		if grid[r][c] != '1' {
			return
		}

		grid[r][c] = '0'
		visit(r-1, c)
		visit(r+1, c)
		visit(r, c-1)
		visit(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				visit(r, c)
			}
		}
	}

	return count
}

func TestNumIslands(t *testing.T) {
	t.Run("grid = example 1", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		expected := 1
		result := numIslands(grid)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("grid = example 2", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}
		expected := 3
		result := numIslands(grid)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("grid = all water", func(t *testing.T) {
		grid := [][]byte{
			{'0', '0', '0'},
			{'0', '0', '0'},
		}
		expected := 0
		result := numIslands(grid)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("grid = single land", func(t *testing.T) {
		grid := [][]byte{
			{'1'},
		}
		expected := 1
		result := numIslands(grid)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("grid = diagonal lands", func(t *testing.T) {
		grid := [][]byte{
			{'1', '0', '1'},
			{'0', '1', '0'},
			{'1', '0', '1'},
		}
		expected := 5
		result := numIslands(grid)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
