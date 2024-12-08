package main

import (
	"fmt"
	"github.com/DennisPing/AOC-2024/utils"
	"log"
	"os"
	"sort"
	"strings"
)

const word1 = "XMAS"
const word2 = "SAMX"
const diagonalSum = 'M' + 'S'

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid [][]rune) {
	count := 0
	count += searchGrid(grid)                   // Regular grid
	count += searchGrid(utils.Transpose(grid))  // Transpose the grid
	count += searchGrid(toDiamond(grid, true))  // Convert grid to diamond of positive diagonals
	count += searchGrid(toDiamond(grid, false)) // Convert grid to diamond of negative diagonals
	fmt.Println(count)
}

func part2(grid [][]rune) {
	count := 0
	n := len(grid)
	m := len(grid[0])

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 'A' {
				found := searchXMAS(grid, i, j)
				if found {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func searchGrid(grid [][]rune) int {
	totalCount := 0
	for _, rawLine := range grid {
		row := string(rawLine)
		totalCount += countSubstr(row, word1)
		totalCount += countSubstr(row, word2)
	}

	return totalCount
}

// Count the number of matching substrings
func countSubstr(row string, substr string) int {
	count := 0
	i := 0
	for i < len(row) {
		j := strings.Index(row[i:], substr)
		if j < 0 {
			break
		}

		count++
		i += j + len(substr)
	}

	return count
}

// Convert a rectangular grid into a diamond grid of diagonals
func toDiamond(grid [][]rune, positiveSlope bool) [][]rune {
	n := len(grid)
	m := len(grid[0])

	// Store diagonals by their indices (row + col)
	diagonals := make(map[int][]rune)

	if positiveSlope {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				diagonals[i-j] = append(diagonals[i-j], grid[i][j])
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				diagonals[i+j] = append(diagonals[i+j], grid[i][j])
			}
		}
	}

	// Collect the diagonals and turn them into slices
	diamond := make([][]rune, len(diagonals))
	keys := make([]int, 0, len(diagonals))
	for k := range diagonals {
		keys = append(keys, k)
	}

	// Sort keys to maintain order
	sort.Ints(keys)
	for i, k := range keys {
		diamond[i] = diagonals[k]
	}

	return diamond
}

// Search for MAS in an X-pattern. Ignore MAM and SAS diagonals.
func searchXMAS(grid [][]rune, n, m int) bool {
	ok1 := grid[n-1][m-1]+grid[n+1][m+1] == diagonalSum
	ok2 := grid[n-1][m+1]+grid[n+1][m-1] == diagonalSum
	return ok1 && ok2
}

func parseInput(fname string) [][]rune {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(content), "\n")
	n := len(lines)
	m := len(lines[0])
	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, m)
	}

	for r, line := range lines {
		for c, char := range line {
			grid[r][c] = char
		}
	}

	return grid
}
