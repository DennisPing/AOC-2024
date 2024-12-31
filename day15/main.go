package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Game struct {
	grid  [][]rune
	moves []rune
	start [2]int
}

type Offset struct {
	dr, dc int
}

var Dirs = map[rune]Offset{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func main() {
	game := parseInput("input.txt")

	//fmt.Println("Starting grid:")
	//for _, x := range game.grid {
	//	fmt.Println(string(x))
	//}

	part1(game)
}

func part1(game Game) {
	grid := game.grid
	moves := game.moves
	start := game.start

	r, c := start[0], start[1]
	finalGrid := walk(grid, r, c, moves)
	sum := sumCoordinates(finalGrid)
	fmt.Println(sum)
}

func walk(grid [][]rune, r, c int, moves []rune) [][]rune {
	for _, move := range moves {
		dir := Dirs[move]
		nr, nc := r+dir.dr, c+dir.dc

		if !inBounds(grid, nr, nc) {
			continue
		}

		switch grid[nr][nc] {
		case '.':
			grid[nr][nc] = '@'
			grid[r][c] = '.'
			r, c = nr, nc
		case 'O':
			ok, newGrid := pushRocks(grid, nr, nc, dir.dr, dir.dc)
			if ok {
				grid = newGrid
				grid[nr][nc] = '@'
				grid[r][c] = '.'
				r, c = nr, nc
			}
		default:
			continue
		}

		//fmt.Println("Move:", string(move))
		//for _, row := range grid {
		//	fmt.Println(string(row))
		//}
		//fmt.Println()
	}

	return grid
}

// Recursive push rocks
func pushRocks(grid [][]rune, r, c, dr, dc int) (bool, [][]rune) {
	nr, nc := r+dr, c+dc
	if !inBounds(grid, nr, nc) {
		return false, grid
	}

	if grid[nr][nc] == '.' {
		// Next position is open, push the rock
		grid[nr][nc] = 'O'
		grid[r][c] = '.'
		return true, grid
	}

	// Recursively check the next rock in a potential chain
	if grid[nr][nc] == 'O' {
		ok, newGrid := pushRocks(grid, nr, nc, dr, dc)
		if ok {
			// Push the current rock
			newGrid[nr][nc] = 'O'
			newGrid[r][c] = '.'
			return true, newGrid
		}
	}

	// Chain in blocked
	return false, grid
}

func inBounds(grid [][]rune, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] != '#'
}

func sumCoordinates(grid [][]rune) int {
	var sum int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'O' {
				sum += (100 * r) + c
			}
		}
	}

	return sum
}

func parseInput(filename string) Game {

	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	parts := strings.Split(string(data), "\n\n")

	var grid [][]rune
	var start [2]int

	for i, line := range strings.Split(parts[0], "\n") {
		j := strings.Index(line, "@")
		if j > 0 {
			start = [2]int{i, j}
		}

		grid = append(grid, []rune(line))
	}

	moves := []rune(parts[1])
	return Game{grid, moves, start}
}
