package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	r, c int
}

type Pair struct {
	p1, p2 Node
}

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid [][]rune) {
	uniqueNodes := make(map[Node]bool)

	nodesMap := findNodes(grid)
	for _, nodes := range nodesMap {
		if len(nodes) < 2 {
			continue // Cannot form any pair
		}
		pairCombinations := getPairCombinations(nodes)
		for _, pair := range pairCombinations {
			extNodes := calcExtendedNodes(grid, pair)
			for _, node := range extNodes {
				uniqueNodes[node] = true
			}
		}
	}

	// printGrid(grid, uniqueNodes)
	fmt.Println(len(uniqueNodes))
}

func part2(grid [][]rune) {
	uniqueNodes := make(map[Node]bool)

	nodesMap := findNodes(grid)
	for _, nodes := range nodesMap {
		if len(nodes) < 2 {
			continue // Cannot form any pair
		}
		pairCombinations := getPairCombinations(nodes)
		for _, pair := range pairCombinations {
			extNodes := calcExtendedNodesRepeated(grid, pair)
			for _, node := range extNodes {
				uniqueNodes[node] = true
			}
		}
	}

	// printGrid(grid, uniqueNodes)
	fmt.Println(len(uniqueNodes))
}

func inBounds(r, c, n, m int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}

func findNodes(grid [][]rune) map[rune][]Node {
	nodes := make(map[rune][]Node)
	for r, row := range grid {
		for c, char := range row {
			if char != '.' {
				nodes[char] = append(nodes[char], Node{r, c})
			}
		}
	}

	return nodes
}

func getPairCombinations(nodes []Node) []Pair {
	var pairs []Pair
	for i, p1 := range nodes {
		for j, p2 := range nodes {
			if i == j {
				continue
			}
			pairs = append(pairs, Pair{p1, p2})
		}
	}

	return pairs
}

func calcExtendedNodes(grid [][]rune, pair Pair) []Node {
	validNodes := make([]Node, 0)

	r1, c1 := pair.p1.r, pair.p1.c
	r2, c2 := pair.p2.r, pair.p2.c

	dr := r2 - r1
	dc := c2 - c1

	ext1 := Node{r1 - dr, c1 - dc}
	ext2 := Node{r2 + dr, c2 + dc}

	if inBounds(ext1.r, ext1.c, len(grid), len(grid[0])) {
		validNodes = append(validNodes, ext1)
	}

	if inBounds(ext2.r, ext2.c, len(grid), len(grid[0])) {
		validNodes = append(validNodes, ext2)
	}

	return validNodes
}

func calcExtendedNodesRepeated(grid [][]rune, pair Pair) []Node {
	validNodes := make([]Node, 0)

	// Append the original nodes because they're now valid
	validNodes = append(validNodes, pair.p1, pair.p2)

	r1, c1 := pair.p1.r, pair.p1.c
	r2, c2 := pair.p2.r, pair.p2.c

	dr := r2 - r1
	dc := c2 - c1

	curR, curC := r1, c1
	for {
		curR -= dr
		curC -= dc
		if inBounds(curR, curC, len(grid), len(grid[0])) {
			validNodes = append(validNodes, Node{curR, curC})
		} else {
			break
		}
	}

	curR, curC = r2, c2
	for {
		curR += dr
		curC += dc
		if inBounds(curR, curC, len(grid), len(grid[0])) {
			validNodes = append(validNodes, Node{curR, curC})
		} else {
			break
		}
	}

	return validNodes
}

func printGrid(grid [][]rune, uniqueNodes map[Node]bool) {
	for r, row := range grid {
		for c, char := range row {
			if _, ok := uniqueNodes[Node{r, c}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func parseInput(fname string) [][]rune {
	content, err := os.ReadFile(fname)
	if err != nil || len(content) == 0 {
		log.Fatalln("unable to read file:", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return grid
}
