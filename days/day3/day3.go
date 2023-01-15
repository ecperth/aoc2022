package day3

import (
	"aoc2022/days"
	file "aoc2022/utils"
	"strconv"
	"unicode"
)

var input = file.ReadInput(3)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	result := 0
	for _, rucksack := range input {
		left := rucksack[:len(rucksack)/2]
		right := rucksack[len(rucksack)/2:]
		foundMatch := false

		for _, i := range left {
			for _, j := range right {
				if i == j {
					if unicode.IsUpper(i) {
						result += int(i - 38)
					} else {
						result += int(i - 96)
					}
					foundMatch = true
					break
				}
			}
			if foundMatch {
				break
			}
		}
	}
	return strconv.Itoa(result)
}

func part2() string {

	result := 0
	for rucksack := 0; rucksack < len(input)-2; rucksack = rucksack + 3 {
		foundMatch := false
		for _, i := range input[rucksack] {
			for _, j := range input[rucksack+1] {
				if i == j {
					for _, k := range input[rucksack+2] {
						if i == k {
							if unicode.IsUpper(i) {
								result += int(i - 38)
							} else {
								result += int(i - 96)
							}
							foundMatch = true
							break
						}
					}
				}
				if foundMatch {
					break
				}
			}
			if foundMatch {
				break
			}
		}
	}

	return strconv.Itoa(result)
}
