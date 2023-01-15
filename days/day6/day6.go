package day6

import (
	"aoc2022/days"
	file "aoc2022/utils"
	"strconv"
)

var input = file.ReadInput(6)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	pos := findFirstUniqueSequence(input[0], 4)
	if pos == -1 {
		return "Not found"
	}
	return strconv.Itoa(pos)
}

func part2() string {

	pos := findFirstUniqueSequence(input[0], 14)
	if pos == -1 {
		return "Not found"
	}
	return strconv.Itoa(pos)
}

//Note part 1 solution worked for part 2 so just broke this out.
/*
findFirstUniqueSequence Finds the position of the first sequence of length [size]
within the sequence [sqnc] containing no repeating characters
*/
func findFirstUniqueSequence(sqnc string, size int) int {
	var buffer []rune
	for pos, c := range sqnc {
		if pos <= size-1 {
			buffer = append(buffer, c)
		} else {
			buffer[pos%size] = c
		}

		if pos >= size-1 {
			isAllUniqueChars := true
			for i, v1 := range buffer {
				if !isAllUniqueChars {
					break
				}
				for _, v2 := range buffer[i+1:] {
					if v1 == v2 {
						isAllUniqueChars = false
						break
					}
				}
			}
			if isAllUniqueChars {
				return pos
			}
		}
	}
	return -1
}
