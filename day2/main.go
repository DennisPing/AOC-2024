package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := parseInput("input.txt")
	part1(data)
	part2(data)
}

func part1(data [][]int) {
	count := 0
	for _, row := range data {
		isSafe := isRowSafe(row)
		if isSafe {
			count++
		}
	}
	fmt.Println(count)
}

func part2(data [][]int) {
	count := 0
	for _, row := range data {
		isSafe := isRowSafe(row)
		if !isSafe {
			isSafe = isFixable(row)
		}
		if isSafe {
			count++
		}
	}
	fmt.Println(count)
}

func isRowSafe(row []int) bool {
	// Choose a direction
	incr := row[1] > row[0]

	for i := 0; i < len(row)-1; i++ {
		x := row[i]
		y := row[i+1]

		var safe bool
		if incr {
			safe = checkPair(x, y)
		} else {
			safe = checkPair(y, x)
		}

		if !safe {
			return false
		}
	}
	return true
}

func checkPair(first, second int) bool {
	return (second-first) >= 1 && (second-first) <= 3
}

func isFixable(row []int) bool {
	for i := 0; i < len(row); i++ {
		// Remove one element and check if OK
		modified := append([]int{}, row[:i]...)
		modified = append(modified, row[i+1:]...)

		if isRowSafe(modified) {
			return true
		}
	}
	return false
}

func parseInput(fname string) [][]int {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	rows := make([][]int, 0)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		row := make([]int, len(fields))
		for i, field := range fields {
			num, _ := strconv.Atoi(field)
			row[i] = num
		}

		rows = append(rows, row)
	}

	return rows
}
