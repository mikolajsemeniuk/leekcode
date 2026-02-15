package leekcode

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/flood-fill
//
// 733. Flood Fill
//
// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
//
// You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].
//
// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
// plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
// Replace the color of all of the aforementioned pixels with color.
//
// Return the modified image after performing the flood fill.
//
// Example 1:
//
// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel
// (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
//
// Example 2:
//
// Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
// Output: [[0,0,0],[0,0,0]]
// Explanation: The starting pixel is already colored 0, so no changes are made to the image.
//
// Constraints:
//
// m == image.length
// n == image[i].length
// 1 <= m, n <= 50
// 0 <= image[i][j], color < 2^16
// 0 <= sr < m
// 0 <= sc < n
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// If the starting pixel already has the target color, no need to do anything
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}

	// Perform DFS to fill all connected pixels
	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, row, col int, originalColor, newColor int) {
	// Check boundaries
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) {
		return
	}

	// Check if current pixel has the original color
	current := image[row][col]
	if current != originalColor {
		return
	}

	// Fill current pixel with new color
	image[row][col] = newColor

	// Recursively fill 4-directionally connected pixels
	dfs(image, row-1, col, originalColor, newColor) // up
	dfs(image, row+1, col, originalColor, newColor) // down
	dfs(image, row, col-1, originalColor, newColor) // left
	dfs(image, row, col+1, originalColor, newColor) // right
}

func TestFloodFill(t *testing.T) {
	t.Run("image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2", func(t *testing.T) {
		image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
		sr := 1
		sc := 1
		color := 2
		expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0", func(t *testing.T) {
		image := [][]int{{0, 0, 0}, {0, 0, 0}}
		sr := 0
		sc := 0
		color := 0
		expected := [][]int{{0, 0, 0}, {0, 0, 0}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[0,0,0],[0,1,1]], sr = 1, sc = 1, color = 1", func(t *testing.T) {
		image := [][]int{{0, 0, 0}, {0, 1, 1}}
		sr := 1
		sc := 1
		color := 1
		expected := [][]int{{0, 0, 0}, {0, 1, 1}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[1]], sr = 0, sc = 0, color = 2", func(t *testing.T) {
		image := [][]int{{1}}
		sr := 0
		sc := 0
		color := 2
		expected := [][]int{{2}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[1,1,1],[1,1,1],[1,1,1]], sr = 1, sc = 1, color = 2", func(t *testing.T) {
		image := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
		sr := 1
		sc := 1
		color := 2
		expected := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[0,0,0],[0,1,0],[0,0,0]], sr = 1, sc = 1, color = 2", func(t *testing.T) {
		image := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
		sr := 1
		sc := 1
		color := 2
		expected := [][]int{{0, 0, 0}, {0, 2, 0}, {0, 0, 0}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[1,1,0],[1,0,1],[0,1,1]], sr = 0, sc = 0, color = 3", func(t *testing.T) {
		image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 1, 1}}
		sr := 0
		sc := 0
		color := 3
		expected := [][]int{{3, 3, 0}, {3, 0, 1}, {0, 1, 1}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[1,2,3],[4,5,6],[7,8,9]], sr = 1, sc = 1, color = 0", func(t *testing.T) {
		image := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		sr := 1
		sc := 1
		color := 0
		expected := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[0,0,0],[0,0,0],[0,0,0]], sr = 1, sc = 1, color = 2", func(t *testing.T) {
		image := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
		sr := 1
		sc := 1
		color := 2
		expected := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("image = [[1,1,1],[1,1,1],[1,1,2]], sr = 2, sc = 2, color = 3", func(t *testing.T) {
		image := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 2}}
		sr := 2
		sc := 2
		color := 3
		expected := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 3}}
		result := floodFill(image, sr, sc, color)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
