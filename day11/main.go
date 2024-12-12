package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair [2]int

var memo = make(map[Pair]int)

func main() {
	stones := parseInput("input.txt")
	part1(stones)
	part2(stones)
}

func part1(stones []int) {
	count := blink1(stones, 25)
	fmt.Println(count)
}

func part2(stones []int) {
	count := blink2(stones, 75)
	fmt.Println(count)
}

func blink1(stones []int, times int) int {
	if times == 0 {
		return len(stones)
	}

	newStones := make([]int, 0)
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if evenDigits(stone) {
			first, second := splitDigits(stone)
			newStones = append(newStones, first, second)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return blink1(newStones, times-1)
}

func blink2(stones []int, times int) int {
	count := 0
	for _, stone := range stones {
		count += scoreStone(stone, times)
	}

	return count
}

func scoreStone(stone int, times int) int {
	if times == 0 {
		return 1 // Score itself
	}

	pair := Pair{stone, times}
	if val, exists := memo[pair]; exists {
		return val
	}

	var result int
	if stone == 0 {
		result = scoreStone(1, times-1)
	} else if evenDigits(stone) {
		first, second := splitDigits(stone)
		result = scoreStone(first, times-1) + scoreStone(second, times-1)
	} else {
		result = scoreStone(stone*2024, times-1)
	}

	// Cache the result
	memo[pair] = result
	return result
}

func evenDigits(stone int) bool {
	return len(strconv.Itoa(stone))%2 == 0
}

func splitDigits(stone int) (int, int) {
	s := strconv.Itoa(stone)
	m := len(s) / 2
	s1 := s[:m]
	s2 := s[m:]

	one, _ := strconv.Atoi(s1)
	two, _ := strconv.Atoi(s2)
	return one, two
}

func parseInput(filename string) []int {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("unable to read file:", err)
	}

	fields := strings.Fields(string(content))
	numbers := make([]int, len(fields))
	for i, v := range fields {
		numbers[i], _ = strconv.Atoi(v)
	}

	return numbers
}
