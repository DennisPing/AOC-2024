package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	R, C, H int
}

var dirs = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func main() {
	grid, trailhead := parseInput("input.txt")
	walkTrail(grid, trailhead)
}

func walkTrail(grid [][]int, trailheads []Point) {
	total1, total2 := 0, 0
	for _, trailhead := range trailheads {
		// Part1
		visited := make(map[Point]bool)
		total1 += dfs(grid, trailhead, 9, visited)

		// Part2
		total2 += dfsNoneVisited(grid, trailhead, 9)
	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func dfs(grid [][]int, cur Point, target int, visited map[Point]bool) int {
	visited[cur] = true
	if cur.H == target {
		return 1
	}

	count := 0
	neighbors := findNeighbors(grid, cur)
	for _, nb := range neighbors {
		if !visited[nb] {
			count += dfs(grid, nb, target, visited)
		}
	}

	return count
}

func dfsNoneVisited(grid [][]int, cur Point, target int) int {
	if cur.H == target {
		return 1
	}

	count := 0
	neighbors := findNeighbors(grid, cur)
	for _, nb := range neighbors {
		count += dfsNoneVisited(grid, nb, target)
	}

	return count
}

func findNeighbors(grid [][]int, cur Point) []Point {
	n := len(grid)
	m := len(grid[0])

	var points []Point
	for _, dir := range dirs {
		nextR := cur.R + dir[0]
		nextC := cur.C + dir[1]
		if inBounds(nextR, nextC, n, m) {
			height := grid[nextR][nextC]
			if cur.H+1 == height {
				points = append(points, Point{nextR, nextC, height})
			}
		}
	}

	return points
}

func inBounds(r, c, n, m int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}

func parseInput(filename string) ([][]int, []Point) {
	content, err := os.ReadFile(filename)
	if err != nil || len(content) == 0 {
		log.Fatalln("unable to read file:", err)
	}

	lines := strings.Split(string(content), "\n")

	var grid [][]int
	var trailheads []Point
	for i, line := range lines {
		row := make([]int, len(line))
		for j, char := range line {
			if char >= '0' && char <= '9' {
				num, _ := strconv.Atoi(string(char))
				if num == 0 {
					trailheads = append(trailheads, Point{i, j, 0})
				}
				row[j] = num
			}
		}

		grid = append(grid, row)
	}

	return grid, trailheads
}
