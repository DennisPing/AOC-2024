package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	set1, set2 := parseInput("input.txt")

	part1(set1, set2)
	part2(set1, set2)
}

func part1(set1, set2 []int) {
	sort.Ints(set1)
	sort.Ints(set2)

	total := 0
	for i := 0; i < len(set1); i++ {
		first := set1[i]
		second := set2[i]

		dist := math.Abs(float64(first - second))
		total += int(dist)
	}

	fmt.Println(total)
}

func part2(set1, set2 []int) {
	sort.Ints(set1)
	sort.Ints(set2)

	freq := make(map[int]int)
	for _, y := range set2 {
		freq[y]++
	}

	total := 0
	for _, x := range set1 {
		total += x * freq[x]
	}

	fmt.Println(total)
}

func parseInput(fname string) ([]int, []int) {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	set1 := make([]int, 0)
	set2 := make([]int, 0)

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		first, _ := strconv.Atoi(fields[0])
		second, _ := strconv.Atoi(fields[1])

		set1 = append(set1, first)
		set2 = append(set2, second)
	}

	return set1, set2
}
