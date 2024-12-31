package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	width  = 101
	height = 103
)

type Position struct {
	x, y int
}

type Velocity struct {
	dx, dy int
}

type Robot struct {
	Pos Position
	Vel Velocity
}

func main() {
	robots := parseInput("input.txt")
	robots2 := slices.Clone(robots) // Because we mutate the robots during simulation

	part1(robots, 100)
	part2(robots2)
}

func part1(robots []Robot, seconds int) {
	for i := 0; i < seconds; i++ {
		for j := range robots {
			robots[j] = updatePos(robots[j])
		}
	}

	counts := countQuadrants(robots)
	sf := safetyFactor(counts)
	fmt.Println(sf)
}

func part2(robots []Robot) {
	for t := 1; ; t++ {
		overlaps := make(map[Position]int)

		for j := range robots {
			robots[j] = updatePos(robots[j])
			overlaps[robots[j].Pos]++
		}

		hasOverlap := false
		for _, freq := range overlaps {
			if freq > 1 {
				hasOverlap = true
				break
			}
		}

		if !hasOverlap {
			fmt.Printf("Seconds: %d\n", t)
			printGrid(robots)
			break
		}
	}
}

func updatePos(robot Robot) Robot {
	robot.Pos.x = (robot.Pos.x + robot.Vel.dx) % width
	if robot.Pos.x < 0 {
		robot.Pos.x += width
	}
	robot.Pos.y = (robot.Pos.y + robot.Vel.dy) % height
	if robot.Pos.y < 0 {
		robot.Pos.y += height
	}
	return robot
}

func countQuadrants(robots []Robot) [4]int {
	var quadrants [4]int
	midW := width / 2
	midH := height / 2

	for _, robot := range robots {
		if robot.Pos.x == midW || robot.Pos.y == midH {
			continue
		}

		if robot.Pos.x < midW && robot.Pos.y < midH {
			quadrants[0]++ // Top left
		} else if robot.Pos.x >= midW && robot.Pos.y < midH {
			quadrants[1]++ // Top right
		} else if robot.Pos.x < midW && robot.Pos.y >= midH {
			quadrants[2]++ // Bot left
		} else if robot.Pos.x >= midW && robot.Pos.y >= midH {
			quadrants[3]++ // Bot right
		}
	}

	return quadrants
}

func safetyFactor(counts [4]int) int {
	sf := 1
	for _, count := range counts {
		sf *= count
	}
	return sf
}

func printGrid(robots []Robot) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, robot := range robots {
		grid[robot.Pos.y][robot.Pos.x] = '#'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func parseInput(filename string) []Robot {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var robots []Robot
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")

		// Parse the Position values
		i := strings.Index(parts[0], "=")
		rawPos := parts[0][i+1:]
		nums := strings.Split(rawPos, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		// Parse the Velocity values
		j := strings.Index(parts[1], "=")
		rawVel := parts[1][j+1:]
		nums = strings.Split(rawVel, ",")
		dx, _ := strconv.Atoi(nums[0])
		dy, _ := strconv.Atoi(nums[1])

		robots = append(robots, Robot{
			Pos: Position{x, y},
			Vel: Velocity{dx, dy}})
	}

	return robots
}
