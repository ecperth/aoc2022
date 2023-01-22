package main

import (
	"aoc2022/days"
	"aoc2022/days/day1"
	"aoc2022/days/day10"
	"aoc2022/days/day11"
	"aoc2022/days/day12"
	"aoc2022/days/day13"
	"aoc2022/days/day14"
	"aoc2022/days/day15"
	"aoc2022/days/day2"
	"aoc2022/days/day3"
	"aoc2022/days/day4"
	"aoc2022/days/day5"
	"aoc2022/days/day6"
	"aoc2022/days/day7"
	"aoc2022/days/day8"
	"aoc2022/days/day9"
	"fmt"
	"os"
	"strconv"
	"time"
)

var solutions = [25]days.Day{
	day1.Solution,
	day2.Solution,
	day3.Solution,
	day4.Solution,
	day5.Solution,
	day6.Solution,
	day7.Solution,
	day8.Solution,
	day9.Solution,
	day10.Solution,
	day11.Solution,
	day12.Solution,
	day13.Solution,
	day14.Solution,
	day15.Solution,
}

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("first argument must be a valid integer"))
	}
	if day > len(solutions) {
		panic(fmt.Errorf("only days 1 -> %d implemented", len(solutions)))
	}

	startTime := time.Now()
	part1 := solutions[day-1].Part1()
	fmt.Printf("Part 1: %s\t Solved in %v\n", part1, time.Now().Sub(startTime))

	startTime = time.Now()
	part2 := solutions[day-1].Part2()
	fmt.Printf("Part 2: %s\t Solved in %v\n", part2, time.Now().Sub(startTime))
}
