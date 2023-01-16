package day8

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
)

var input = utils.ReadInputAsBytes(8)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	//edges
	xMax := len(input[0])
	yMax := len(input)
	visibleTreeCount := 2*xMax + 2*(yMax-2)

	//inner
	for x := 1; x < xMax-1; x++ {
		for y := 1; y < yMax-1; y++ {
			height := input[y][x]
			//left
			if isGreaterThanAllElements(height, input[y][:x]) {
				visibleTreeCount++
				continue
			}
			//right
			if isGreaterThanAllElements(height, input[y][x+1:]) {
				visibleTreeCount++
				continue
			}
			//up
			if isGreaterThanAllElements(height, getNthElementOfEachArray(x, input[:y])) {
				visibleTreeCount++
				continue
			}
			//down
			if isGreaterThanAllElements(height, getNthElementOfEachArray(x, input[y+1:])) {
				visibleTreeCount++
				continue
			}
		}
	}

	return strconv.Itoa(visibleTreeCount)
}

func part2() string {

	//edges
	xMax := len(input[0])
	yMax := len(input)

	topScore := 0
	//can ignore edges as "If a tree is right on the edge, at least one of its viewing distances will be zero."
	for x := 1; x < xMax-1; x++ {
		for y := 1; y < yMax-1; y++ {
			height := input[y][x]
			scenicScore :=
				findViewingDistance(height, input[y][:x], true) *
					findViewingDistance(height, input[y][x+1:], false) *
					findViewingDistance(height, getNthElementOfEachArray(x, input[:y]), true) *
					findViewingDistance(height, getNthElementOfEachArray(x, input[y+1:]), false)
			if scenicScore > topScore {
				topScore = scenicScore
			}
		}
	}
	return strconv.Itoa(topScore)
}

func isGreaterThanAllElements(v byte, elements []byte) bool {
	for i := 0; i < len(elements); i++ {
		if v <= elements[i] {
			return false
		}
	}
	return true
}

func getNthElementOfEachArray(n int, elements [][]byte) []byte {
	result := make([]byte, len(elements))
	for i, v := range elements {
		result[i] = v[n]
	}
	return result
}

// if reverse, then check from top index of the slice
func findViewingDistance(height byte, trees []byte, reverse bool) int {
	offset := 0
	if reverse {
		offset = len(trees) - 1
	}

	for i := 0; i < len(trees); i++ {
		if height <= trees[utils.Abs(i-offset)] {
			return i + 1
		}
	}
	return len(trees)
}
