package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	expr := parseInput("input.txt")
	part1(expr)
	part2(expr)
}

func part1(expr string) {
	// Thanks ChatGPT :)
	// `mul\(` matches the literal "mul(".
	// `\d+` matches one or more digits.
	// `,` matches the literal comma separating the numbers.
	// `\d+` matches one or more digits (second number).
	// `\)` matches the literal closing parenthesis.
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(expr, -1)
	total := 0

	for _, match := range matches {
		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil {
			total += x * y
		}
	}
	fmt.Println(total)
}

func part2(expr string) {
	// Thanks OpenAI :)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(expr, -1)
	enabled := true
	total := 0

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if match[1] != "" && match[2] != "" && enabled {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				total += x * y
			}
		}
	}
	fmt.Println(total)
}

func parseInput(fname string) string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	return string(content)
}
