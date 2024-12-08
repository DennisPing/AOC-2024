package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Vector struct {
	r, c   int
	dr, dc int
}

var directions = []Vector{
	{dr: -1, dc: 0}, // Up
	{dr: 0, dc: 1},  // Right
	{dr: 1, dc: 0},  // Down
	{dr: 0, dc: -1}, // Left
}

func main() {
	grid, start := parseInput("input.txt")
	part1(grid, start)
	part2(grid, start)
}

func part1(grid [][]rune, start Vector) {
	visited := walk(grid, start)
	fmt.Println(len(visited))
}

func part2(grid [][]rune, start Vector) {
	visited := walk(grid, start)
	cycles := countPossibleCycles(grid, start, visited)
	fmt.Println(cycles)
}

func walk(grid [][]rune, start Vector) map[Vector]bool {
	visited := make(map[Vector]bool)
	n := len(grid)
	m := len(grid[0])
	cur := start

	for !atBoundary(n, m, cur) {
		// Get current char
		char := grid[cur.r][cur.c]

		// Get next position
		p := cur.r + cur.dr
		q := cur.c + cur.dc

		if grid[p][q] == '#' {
			// Rotate clockwise in place
			newDir := rotateClockwise(cur)
			cur.dr = newDir.dr
			cur.dc = newDir.dc
			grid[cur.r][cur.c] = char
		} else {
			// Move char to next position
			grid[p][q] = char
			grid[cur.r][cur.c] = '.'

			cur.r, cur.c = p, q
			visited[Vector{r: cur.r, c: cur.c}] = true // Only hash [r,c] and not [dr,dc] part
		}
	}

	return visited
}

func countPossibleCycles(grid [][]rune, start Vector, visited map[Vector]bool) int {
	cycles := 0
	n := len(grid)
	m := len(grid[0])

	// For each visited Vector, treat it as a hypothetical obstacle
	for obst := range visited {
		// Run a fresh scenario
		tempVisited := make(map[Vector]bool)
		cur := start

		for !atBoundary(n, m, cur) {
			if tempVisited[cur] {
				cycles++
				break
			}
			tempVisited[cur] = true

			// Get next position
			p := cur.r + cur.dr
			q := cur.c + cur.dc

			if (p == obst.r && q == obst.c) || grid[p][q] == '#' {
				// Rotate clockwise in place
				newDir := rotateClockwise(cur)
				cur.dr = newDir.dr
				cur.dc = newDir.dc
			} else {
				// Move cur to next position
				cur.r, cur.c = p, q
			}
		}
	}

	return cycles
}

func atBoundary(n, m int, v Vector) bool {
	return v.r == 0 || v.c == 0 || v.r == n-1 || v.c == m-1
}

func rotateClockwise(cur Vector) Vector {
	for i, d := range directions {
		if d.dr == cur.dr && d.dc == cur.dc {
			return directions[(i+1)%4]
		}
	}
	return cur
}

// Returns the 2d grid, starting point, and 2d list of obstacles
func parseInput(fname string) ([][]rune, Vector) {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	n, m := len(lines), len(lines[0])
	grid := make([][]rune, n)
	start := Vector{}

	for r, line := range lines {
		grid[r] = make([]rune, m)
		for c, char := range line {
			grid[r][c] = char
			if char == '^' {
				start = Vector{r: r, c: c, dr: -1, dc: 0}
			}
		}
	}
	return grid, start
}
