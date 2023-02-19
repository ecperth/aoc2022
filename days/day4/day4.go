package day4

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
)

var input = utils.ReadInputAsStrings(4)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func toInts(sections []string) (start int, end int) {
	start = utils.AtoiUnsafe(sections[0])
	end = utils.AtoiUnsafe(sections[1])
	return
}

func part1() string {

	result := 0
	for _, assignment := range input {
		var a1Start, a1End, a2Start, a2End int
		utils.SscanfUnsafe(assignment, "%d-%d,%d-%d", &a1Start, &a1End, &a2Start, &a2End)

		if (a1Start <= a2Start && a1End >= a2End) || (a2Start <= a1Start && a2End >= a1End) {
			result++
		}
	}

	return strconv.Itoa(result)
}

func part2() string {

	result := 0
	for _, assignment := range input {
		var a1Start, a1End, a2Start, a2End int
		utils.SscanfUnsafe(assignment, "%d-%d,%d-%d", &a1Start, &a1End, &a2Start, &a2End)

		if (a1Start <= a2Start && a1End >= a2Start) || (a2Start <= a1Start && a2End >= a1Start) {
			result++
		}
	}

	return strconv.Itoa(result)
}
