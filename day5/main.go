package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates := parseInput("input.txt")
	valids, invalids := checkUpdates(rules, updates)
	part1(valids)
	part2(invalids, rules)
}

func checkUpdates(rules map[int][]int, updates [][]int) (valids [][]int, invalids [][]int) {
	for _, row := range updates {
		ok := checkRow(row, rules)
		if ok {
			valids = append(valids, row)
		} else {
			invalids = append(invalids, row)
		}
	}

	return
}

func part1(valids [][]int) {
	total := 0
	for _, v := range valids {
		total += median(v)
	}
	fmt.Println(total)
}

func part2(invalids [][]int, rules map[int][]int) {
	for i, row := range invalids {
		invalids[i] = fixRow(row, rules)
	}

	total := 0
	for _, iv := range invalids {
		total += median(iv)
	}
	fmt.Println(total)
}

func checkRow(row []int, rules map[int][]int) bool {
	visited := make(map[int]bool)

	for _, num := range row {
		visited[num] = true

		ruleList, ok := rules[num]
		if !ok {
			continue // No rule for this number, move on
		}

		for _, q := range ruleList {
			if visited[q] {
				return false // Row is invalid
			}
		}
	}

	return true
}

func fixRow(row []int, rules map[int][]int) []int {
	// Big while loop. Modify the row in place.
	i := 0
	visited := make(map[int]int)
	for i < len(row) {
		num := row[i]
		visited[num] = i

		ruleList, ok := rules[num]
		if !ok {
			i++
			continue // No rule for this number, move on
		}

		swapped := false
		for _, q := range ruleList {
			if j, ok := visited[q]; ok {
				row[i], row[j] = row[j], row[i]
				swapped = true
				// Clear visited because we just swapped and need to re-check
				for k := range visited {
					delete(visited, k)
				}
				break
			}
		}

		if swapped {
			i = 0
		} else {
			i++
		}
	}
	return row
}

func median(row []int) int {
	mid := len(row) / 2
	return row[mid]
}

func parseInput(fname string) (map[int][]int, [][]int) {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	sections := strings.Split(string(content), "\n\n")
	if len(sections) != 2 {
		log.Fatalf("expected 2 sections, got %d", len(sections))
	}

	// Parse the rules section
	rules := make(map[int][]int)
	lines := strings.Split(sections[0], "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		rules[x] = append(rules[x], y)
	}

	// Parse the updates section
	updates := make([][]int, 0)
	lines = strings.Split(sections[1], "\n")
	for _, line := range lines {
		row := make([]int, 0)
		fields := strings.Split(line, ",")
		for _, word := range fields {
			num, _ := strconv.Atoi(word)
			row = append(row, num)
		}
		updates = append(updates, row)
	}

	return rules, updates
}
