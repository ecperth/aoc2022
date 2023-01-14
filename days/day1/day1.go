package day1

import (
	"aoc2022/days"
	"aoc2022/utils"
	"sort"
	"strconv"
)

var input = utils.ReadInput(1)
var Day1 = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() int {

	mostCalories := 0
	calorieCounter := 0
	for i, line := range input {
		if len(line) == 0 || i == len(input) {
			if calorieCounter > mostCalories {
				mostCalories = calorieCounter
			}
			calorieCounter = 0
		} else {
			calories, _ := strconv.Atoi(line)
			calorieCounter += calories
		}
	}

	return mostCalories
}

func part2() int {

	topThree := []int{0, 0, 0}
	calorieCounter := 0

	for i, line := range input {
		if len(line) == 0 || i == len(input) {
			if calorieCounter > topThree[0] {
				topThree[0] = calorieCounter
				sort.Ints(topThree)
			}
			calorieCounter = 0
		} else {
			calories, _ := strconv.Atoi(line)
			calorieCounter += calories
		}
	}

	result := 0
	for _, calories := range topThree {
		result += calories
	}

	return result
}
