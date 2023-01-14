package day4

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInput(4)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func toInts(sections []string) (int, int) {
	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])
	return start, end
}

func part1() string {

	result := 0
	for _, assignment := range input {
		a1Start, a1End := toInts(strings.Split(strings.Split(assignment, ",")[0], "-"))
		a2Start, a2End := toInts(strings.Split(strings.Split(assignment, ",")[1], "-"))

		if (a1Start <= a2Start && a1End >= a2End) || (a2Start <= a1Start && a2End >= a1End) {
			result++
		}
	}

	return strconv.Itoa(result)
}

func part2() string {

	result := 0
	for _, assignment := range input {
		a1Start, a1End := toInts(strings.Split(strings.Split(assignment, ",")[0], "-"))
		a2Start, a2End := toInts(strings.Split(strings.Split(assignment, ",")[1], "-"))

		if (a1Start <= a2Start && a1End >= a2Start) || (a2Start <= a1Start && a2End >= a1Start) {
			result++
		}
	}

	return strconv.Itoa(result)
}
