package utils

// Transpose a 2D slice of any type
func Transpose[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return [][]T{}
	}
	n := len(grid)
	m := len(grid[0])
	transposed := make([][]T, m)
	for i := range transposed {
		transposed[i] = make([]T, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transposed[j][i] = grid[i][j]
		}
	}

	return transposed
}

// Rotate90 rotates a 2D slice 90 degrees clockwise.
func Rotate90[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return [][]T{}
	}
	n := len(grid)
	m := len(grid[0])
	rotated := make([][]T, m)
	for i := range rotated {
		rotated[i] = make([]T, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotated[j][n-1-i] = grid[i][j]
		}
	}

	return rotated
}
