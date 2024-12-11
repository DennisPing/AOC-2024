package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// -- Disk Block abstraction --

type Block struct {
	Data []int
}

// Size returns the data size of the block.
func (b *Block) Size() int {
	return len(b.Data)
}

// Cap returns the fixed capacity of the block.
func (b *Block) Cap() int {
	return cap(b.Data)
}

// Space returns the remaining space of the block.
func (b *Block) Space() int {
	return b.Cap() - b.Size()
}

// -- Solution --

func main() {
	numbers := parseInput("input.txt")
	part1(toBlocks(numbers))
	part2(toBlocks(numbers))
}

func part1(blocks []Block) {
	blocks = moveBlocks(blocks)
	fmt.Println(checksum(blocks))
}

func part2(blocks []Block) {
	blocks = moveFile(blocks)
	fmt.Println(checksum(blocks))
}

func toBlocks(numbers []int) []Block {
	var blocks []Block
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			// Data block
			size := numbers[i]
			value := i / 2
			data := slices.Repeat([]int{value}, size)
			blocks = append(blocks, Block{data})
		} else {
			// Empty block
			capacity := numbers[i]
			data := make([]int, 0, capacity)
			blocks = append(blocks, Block{data})
		}
	}

	return blocks
}

func moveBlocks(blocks []Block) []Block {
	head := 1

	// Find the last data block
	tail := len(blocks) - 1
	for tail >= 0 && blocks[tail].Size() == 0 {
		tail--
	}

	// Move blocks from tail to head
	for tail > head {
		if blocks[head].Space() == 0 {
			head += 2
			continue
		}
		if blocks[tail].Size() == 0 {
			tail -= 2
			continue
		}

		moveData(&blocks[tail], &blocks[head])
	}

	return blocks
}

func moveFile(blocks []Block) []Block {
	head := 1

	// Find the last data block
	tail := len(blocks) - 1
	for tail >= 0 && blocks[tail].Size() == 0 {
		tail--
	}

	for tail > head {
		if blocks[head].Space() == 0 {
			head += 2
			continue
		}
		if blocks[tail].Size() == 0 {
			tail -= 2
			continue
		}

		// Search backwards
		for j := tail; j > head; j -= 2 {
			moved := false

			// Search forwards
			for i := head; i < tail; i += 2 {
				if blocks[j].Size() <= blocks[i].Space() {
					moveAllData(&blocks[j], &blocks[i])
					moved = true
					break
				}
			}

			if moved {
				tail -= 2
				break
			}
		}
	}

	return blocks
}

func moveData(from *Block, to *Block) {
	end := from.Size() - 1
	to.Data = append(to.Data, from.Data[end])
	from.Data = from.Data[:end]
}

func moveAllData(from *Block, to *Block) {
	to.Data = append(to.Data, from.Data...)
	from.Data = from.Data[:0] // Clear data while keeping capacity
}

func checksum(blocks []Block) int {
	var sum, i int
	for _, block := range blocks {
		for _, val := range block.Data {
			prod := i * val
			sum += prod
			i++
		}
		// Make sure to increment over any unused space
		i += block.Space()
	}

	return sum
}

func parseInput(filename string) []int {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("unable to read file:", err)
	}

	line := strings.TrimSpace(string(content))

	numbers := make([]int, len(line))
	for i, char := range line {
		num, _ := strconv.Atoi(string(char))
		numbers[i] = num
	}

	return numbers
}
