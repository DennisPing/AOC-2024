package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	R, C int
}

var dirs = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

var cornerPatterns = [][][2]int{
	{
		// Top right corner
		{-1, 0},
		{0, 1},
	},
	{
		// Bot right corner
		{1, 0},
		{0, 1},
	},
	{
		// Bot left corner
		{1, 0},
		{0, -1},
	},
	{
		// Top left corner
		{-1, 0},
		{0, -1},
	},
}

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid [][]rune) {
	total := 0
	visited := make(map[Point]bool)
	cur := Point{0, 0}
	for i, row := range grid {
		for j, char := range row {
			cur = Point{i, j}
			if !visited[cur] {
				region := make([]Point, 0)
				region = floodFillDfs(grid, char, cur, region, visited)
				total += getArea(region) * getPerimeter(grid, region)
			}
		}
	}

	fmt.Println(total)
}

func part2(grid [][]rune) {
	total := 0
	visited := make(map[Point]bool)
	cur := Point{0, 0}
	for i, row := range grid {
		for j, char := range row {
			cur = Point{i, j}
			if !visited[cur] {
				region := make([]Point, 0)
				region = floodFillDfs(grid, char, cur, region, visited)
				total += getArea(region) * getSides(grid, region)
			}
		}
	}

	fmt.Println(total)
}

func floodFillDfs(grid [][]rune, char rune, cur Point, region []Point, visited map[Point]bool) []Point {
	if visited[cur] || !inBounds(cur.R, cur.C, len(grid), len(grid[0])) || grid[cur.R][cur.C] != char {
		return region
	}

	visited[cur] = true
	region = append(region, cur)

	neighbors := getNeighbors(grid, cur)
	for _, nb := range neighbors {
		region = floodFillDfs(grid, char, nb, region, visited)
	}

	return region
}

func getNeighbors(grid [][]rune, cur Point) []Point {
	var neighbors []Point // oob = out of bounds
	for _, dir := range dirs {
		nr := cur.R + dir[0]
		nc := cur.C + dir[1]
		if inBounds(nr, nc, len(grid), len(grid[0])) {
			neighbors = append(neighbors, Point{nr, nc})
		}
	}

	return neighbors
}

func getOutOfBounds(grid [][]rune, cur Point) []Point {
	var oobs []Point // oob = out of bounds
	for _, dir := range dirs {
		nr := cur.R + dir[0]
		nc := cur.C + dir[1]
		if !inBounds(nr, nc, len(grid), len(grid[0])) {
			oobs = append(oobs, Point{nr, nc})
		}
	}

	return oobs
}

func inBounds(r, c, n, m int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}

func getArea(region []Point) int {
	return len(region)
}

func getPerimeter(grid [][]rune, region []Point) int {
	perim := 0
	for _, point := range region {
		neighbors := getNeighbors(grid, point)
		for _, nb := range neighbors {
			if grid[point.R][point.C] != grid[nb.R][nb.C] {
				perim++
			}
		}
		oobs := getOutOfBounds(grid, point)
		perim += len(oobs)
	}

	return perim
}

func getSides(grid [][]rune, region []Point) int {
	sides := 0
	for _, point := range region {
		count := countCorners(grid, point)
		sides += count
	}

	return sides
}

func countCorners(grid [][]rune, cur Point) int {
	char := grid[cur.R][cur.C]
	count := 0

	for _, pattern := range cornerPatterns {
		// Directional offsets
		dirA := pattern[0]
		dirB := pattern[1]

		// New (row, col)
		rA, cA := cur.R+dirA[0], cur.C+dirA[1]
		rB, cB := cur.R+dirB[0], cur.C+dirB[1]
		rD, cD := rA+dirB[0], cA+dirB[1] // Diagonal cell

		insideA := inBounds(rA, cA, len(grid), len(grid[0])) && grid[rA][cA] == char
		insideB := inBounds(rB, cB, len(grid), len(grid[0])) && grid[rB][cB] == char
		insideD := inBounds(rD, cD, len(grid), len(grid[0])) && grid[rD][cD] == char

		// A cell is a boundary if it's out-of-bounds or different char
		boundaryA := !insideA
		boundaryB := !insideB
		boundaryD := !insideD

		// Convex corner condition:
		// Both A and B are boundaries
		if boundaryA && boundaryB {
			count++
			continue
		}

		// Concave corner condition:
		// A and B are inside, but the diagonal is a boundary
		if insideA && insideB && boundaryD {
			count++
		}
	}

	return count
}

func parseInput(filename string) [][]rune {
	content, err := os.ReadFile(filename)
	if err != nil || len(content) == 0 {
		log.Fatalln("unable to read file:", err)
	}

	var grid [][]rune
	for _, line := range strings.Split(string(content), "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}
