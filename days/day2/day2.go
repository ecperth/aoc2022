package day2

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInput(2)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	shapes := map[string]int{
		"A": 0,
		"X": 0,
		"B": 1,
		"Y": 1,
		"C": 2,
		"Z": 2,
	}

	totalScore := 0
	for _, round := range input {
		moves := strings.Split(round, " ")

		roundScore := shapes[moves[1]] + 1
		result := (shapes[moves[1]] - shapes[moves[0]] + 3) % 3

		if result == 0 {
			roundScore += 3
		} else if result == 1 {
			roundScore += 6
		}
		totalScore += roundScore
	}

	return strconv.Itoa(totalScore)
}

func part2() string {

	shapes := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}

	totalScore := 0
	for _, round := range input {
		moves := strings.Split(round, " ")

		switch moves[1] {
		case "X":
			totalScore += (shapes[moves[0]]+2)%3 + 1
		case "Y":
			totalScore += 3
			totalScore += shapes[moves[0]] + 1
		case "Z":
			totalScore += 6
			totalScore += (shapes[moves[0]]+1)%3 + 1
		}
	}

	return strconv.Itoa(totalScore)
}
