package day10

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(10)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	rv, c := 1, 0
	sigStrengths := 0
	for _, inst := range input {
		tokens := strings.Split(inst, " ")
		if (c-19)%40 == 0 {
			sigStrengths += rv * (c + 1)
		}
		c++
		if tokens[0] == "addx" {
			if (c-19)%40 == 0 {
				sigStrengths += rv * (c + 1)
			}
			c++
			v, _ := strconv.Atoi(tokens[1])
			rv += v
		}
	}

	return strconv.Itoa(sigStrengths)
}

func part2() string {
	rv, c := 1, 0
	display := []byte{'\n'}
	for _, inst := range input {
		tokens := strings.Split(inst, " ")
		if c%40 >= rv-1 && c%40 <= rv+1 {
			display = append(display, byte('#'))
		} else {
			display = append(display, byte('.'))
		}
		c++
		if tokens[0] == "addx" {
			if c%40 == 0 {
				display = append(display, byte('\n'))
			}
			if c%40 >= rv-1 && c%40 <= rv+1 {
				display = append(display, byte('#'))
			} else {
				display = append(display, byte('.'))
			}
			c++
			v, _ := strconv.Atoi(tokens[1])
			rv += v
		}
		if c%40 == 0 {
			display = append(display, byte('\n'))
		}
	}

	return string(display)
}
