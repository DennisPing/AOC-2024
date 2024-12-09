package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func main() {
	equations := parseInput("input.txt")
	part1(equations)
	part2(equations)
}

func part1(equations []Equation) {
	sum := 0
	for _, eq := range equations {
		if backtrack1(eq.Numbers, 1, eq.Numbers[0], eq.TestValue) {
			sum += eq.TestValue
		}
	}

	fmt.Println(sum)
}

func part2(equations []Equation) {
	sum := 0
	for _, eq := range equations {
		if backtrack2(eq.Numbers, 1, eq.Numbers[0], eq.TestValue) {
			sum += eq.TestValue
		}
	}

	fmt.Println(sum)
}

func backtrack1(nums []int, idx, currentVal, target int) bool {
	// Base case 1: If the current value is already greater than the target, return false
	if currentVal > target {
		return false
	}

	// Base case 2: If we've used up all the numbers, check if the current value is equal to target
	if idx == len(nums) {
		return currentVal == target
	}

	nextNum := nums[idx]

	// Try addition
	if backtrack1(nums, idx+1, currentVal+nextNum, target) {
		return true
	}

	// Try multiplication
	if backtrack1(nums, idx+1, currentVal*nextNum, target) {
		return true
	}

	return false
}

func backtrack2(nums []int, idx, currentVal, target int) bool {
	// Base case 1: If the current value is already greater than the target, return false
	if currentVal > target {
		return false
	}

	// Base case 2: If we've used up all the numbers, check if the current value is equal to target
	if idx == len(nums) {
		return currentVal == target
	}

	nextNum := nums[idx]

	// Try addition
	if backtrack2(nums, idx+1, currentVal+nextNum, target) {
		return true
	}

	// Try multiplication
	if backtrack2(nums, idx+1, currentVal*nextNum, target) {
		return true
	}

	// Try concat
	concatVal := concatNumbers(currentVal, nextNum)
	if backtrack2(nums, idx+1, concatVal, target) {
		return true
	}

	return false
}

func concatNumbers(a, b int) int {
	// Count digits in b
	digits := 1
	temp := b
	for temp >= 10 {
		temp /= 10
		digits++
	}

	return a*pow10(digits) + b
}

func pow10(n int) int {
	val := 1
	for i := 0; i < n; i++ {
		val *= 10
	}
	return val
}

func parseInput(fname string) []Equation {
	content, err := os.ReadFile(fname)
	if err != nil || len(content) == 0 {
		log.Fatalln("unable to read file:", err)
	}

	var equations []Equation

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		var numbers []int

		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])

		for _, rawNum := range strings.Split(parts[1], " ") {
			num, _ := strconv.Atoi(rawNum)
			numbers = append(numbers, num)
		}

		equations = append(equations, Equation{TestValue: testValue, Numbers: numbers})
	}

	return equations
}
