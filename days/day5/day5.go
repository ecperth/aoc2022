package day5

import (
	"aoc2022/days"
	file "aoc2022/utils"
	"fmt"
)

var input = file.ReadInput(5)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

const (
	stackCount = 9
)

// pop n items off slice and retain order
func pop(stack *[]rune, n int) []rune {
	v := (*stack)[len(*stack)-n:]
	*stack = (*stack)[:len(*stack)-n]
	return v
}

// push n items to slice and retain order
func push(stack *[]rune, v []rune) {
	*stack = append(*stack, v...)
}

func part1() string {

	var blankLineIndex int
	for i, line := range input {
		if line == "" {
			blankLineIndex = i
			break
		}
	}

	//init stacks
	var supplyStacks [stackCount][]rune
	for line := blankLineIndex - 2; line >= 0; line-- {
		row := input[line]
		cursor := 1

		for s := 0; s < stackCount; s++ {
			if row[cursor] != 32 {
				supplyStacks[s] = append(supplyStacks[s], rune(row[cursor]))
			}
			cursor = cursor + 4
			if cursor > len(row) {
				break
			}
		}
	}

	//process commands
	for line := blankLineIndex + 1; line < len(input); line++ {
		var count, src, dest int
		_, err := fmt.Sscanf(input[line], "move %d from %d to %d", &count, &src, &dest)
		if err != nil {
			panic(fmt.Errorf("could not parse instruction om line %d", line))
		}

		for i := 1; i <= count; i++ {
			push(&supplyStacks[dest-1], pop(&supplyStacks[src-1], 1))
		}
	}

	//results
	var result []rune
	for s := 0; s < stackCount; s++ {
		if len(supplyStacks[s]) >= 1 {
			result = append(result, supplyStacks[s][len(supplyStacks[s])-1])
		}

	}
	return string(result)
}

func part2() string {

	var blankLineIndex int
	for i, line := range input {
		if line == "" {
			blankLineIndex = i
			break
		}
	}

	//init stacks
	var supplyStacks [stackCount][]rune
	for line := blankLineIndex - 2; line >= 0; line-- {
		row := input[line]
		cursor := 1

		for s := 0; s < stackCount; s++ {
			if row[cursor] != 32 {
				supplyStacks[s] = append(supplyStacks[s], rune(row[cursor]))
			}
			cursor = cursor + 4
			if cursor > len(row) {
				break
			}
		}
	}

	//process commands
	for line := blankLineIndex + 1; line < len(input); line++ {
		var count, src, dest int
		_, err := fmt.Sscanf(input[line], "move %d from %d to %d", &count, &src, &dest)
		if err != nil {
			panic(fmt.Errorf("could not parse instruction om line %d", line))
		}

		push(&supplyStacks[dest-1], pop(&supplyStacks[src-1], count))
	}

	//results
	var result []rune
	for s := 0; s < stackCount; s++ {
		if len(supplyStacks[s]) >= 1 {
			result = append(result, supplyStacks[s][len(supplyStacks[s])-1])
		}

	}
	return string(result)
}
