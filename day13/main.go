package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	buttonA Button
	buttonB Button
	prize   [2]int
}

type Button struct {
	x    int
	y    int
	cost int
}

func main() {
	games := parseInput("input.txt")
	part1(games)
	part2(games)
}

func part1(games []Game) {
	total := 0
	for _, game := range games {
		cost := solveGame(game)
		if cost > 0 {
			total += cost
		}
	}

	fmt.Println(total)
}

func part2(game []Game) {
	total := 0
	for _, game := range game {
		game.prize[0] += 10000000000000
		game.prize[1] += 10000000000000
		cost := solveGameScaled(game)
		if cost > 0 {
			total += cost
		}
	}

	fmt.Println(total)
}

func solveGame(game Game) int {
	A := game.buttonA
	B := game.buttonB
	Px := game.prize[0]
	Py := game.prize[1]

	minCost := math.MaxInt
	solvable := false

	// Check if Px and Py actually have GCD with button A and button B
	if Px%gcd(A.x, B.x) != 0 || Py%gcd(A.y, B.y) != 0 {
		return -1
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			cost := i*A.cost + j*B.cost
			if cost >= minCost {
				break
			}

			// Calculate the new claw position
			xPos := i*A.x + j*B.x
			yPos := i*A.y + j*B.y

			if xPos == Px && yPos == Py {
				if cost < minCost {
					minCost = cost
					solvable = true
				}
			}
		}
	}

	if solvable {
		return minCost
	}

	return -1
}

func solveGameScaled(game Game) int {
	A := game.buttonA
	B := game.buttonB
	Px := game.prize[0]
	Py := game.prize[1]

	/*
		System of equations
		Px = Ax * i + Bx * j
		Py = Ay * i + By * j
		where i = number of A presses and j = number of B presses

		j = (Px - Ax * i) / Bx
		j = (Py - Ay * i) / By

		i = (Px * By - Py * Bx) / (Ax * By - Ay * Bx)
		j = (Py * Ax - Px * Ay) / (Ax * By - Ay * Bx)
	*/

	i := float64(Px*B.y-Py*B.x) / float64(A.x*B.y-A.y*B.x)
	j := float64(Py*A.x-Px*A.y) / float64(A.x*B.y-A.y*B.x)

	if i == math.Trunc(i) && j == math.Trunc(j) {
		return A.cost*int(i) + B.cost*int(j)
	}

	return 0
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func parseInput(filename string) []Game {
	content, err := os.ReadFile(filename)
	if err != nil || len(content) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var games []Game

	sections := strings.Split(string(content), "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		first := lines[0]
		second := lines[1]
		third := lines[2]

		buttonA := parseButtonLine(first, 3)
		buttonB := parseButtonLine(second, 1)
		prize := parsePrizeLine(third)

		game := Game{buttonA, buttonB, prize}
		games = append(games, game)
	}

	return games
}

func parseButtonLine(line string, cost int) Button {
	xStr := strings.Index(line, "X+") + 2
	xEnd := strings.Index(line[xStr:], ",")
	yStr := strings.Index(line, "Y+") + 2

	dx, _ := strconv.Atoi(line[xStr : xStr+xEnd])
	dy, _ := strconv.Atoi(line[yStr:])

	return Button{dx, dy, cost}
}

func parsePrizeLine(line string) [2]int {
	xStr := strings.Index(line, "X=") + 2
	xEnd := strings.Index(line[xStr:], ",")
	yStr := strings.Index(line, "Y=") + 2

	x, _ := strconv.Atoi(line[xStr : xStr+xEnd])
	y, _ := strconv.Atoi(line[yStr:])

	return [2]int{x, y}
}
