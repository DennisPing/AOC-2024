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

func main() {
	grid := parseInput("input.txt")
	part1(grid)
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
				total += area(region) * perimeter(grid, region)
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

	neighbors, _ := getNeighbors(grid, cur)
	for _, nb := range neighbors {
		region = floodFillDfs(grid, char, nb, region, visited)
	}

	return region
}

func getNeighbors(grid [][]rune, cur Point) ([]Point, []Point) {
	var neighbors, oobs []Point // oob = out of bounds
	for _, dir := range dirs {
		nextR := cur.R + dir[0]
		nextC := cur.C + dir[1]
		if inBounds(nextR, nextC, len(grid), len(grid[0])) {
			neighbors = append(neighbors, Point{nextR, nextC})
		} else {
			oobs = append(oobs, Point{nextR, nextC})
		}
	}

	return neighbors, oobs
}

func inBounds(r, c, n, m int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}

func area(region []Point) int {
	return len(region)
}

func perimeter(grid [][]rune, region []Point) int {
	count := 0
	for _, point := range region {
		neighbors, oobs := getNeighbors(grid, point)
		for _, nb := range neighbors {
			if grid[point.R][point.C] != grid[nb.R][nb.C] {
				count++
			}
		}
		count += len(oobs)
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
