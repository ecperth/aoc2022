package main

import (
	"aoc2022/days"
	"aoc2022/days/day1"
	"aoc2022/days/day2"
	"aoc2022/days/day3"
	"fmt"
	"os"
	"strconv"
	"time"
)

var solutions = [25]days.Day{
	day1.Day1,
	day2.Day2,
	day3.Day3,
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
	fmt.Printf("Part 1: %d\t Solved in %v\n", part1, time.Now().Sub(startTime))

	startTime = time.Now()
	part2 := solutions[day-1].Part2()
	fmt.Printf("Part 2: %d\t Solved in %v\n", part2, time.Now().Sub(startTime))
}
