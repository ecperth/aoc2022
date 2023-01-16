package day9

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(9)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

var empty struct{}

type coordinate struct {
	x int
	y int
}

func part1() string {

	hPos, tPos, dest := &coordinate{x: 0, y: 0}, &coordinate{x: 0, y: 0}, &coordinate{x: 0, y: 0}
	visitedTpos := map[coordinate]struct{}{
		coordinate{tPos.x, tPos.y}: empty,
	}

	for _, move := range input {
		m := strings.Split(move, " ")
		dir := m[0]
		mag, _ := strconv.Atoi(m[1])

		switch dir {
		case "L":
			dest = &coordinate{hPos.x - mag, hPos.y}
		case "R":
			dest = &coordinate{hPos.x + mag, hPos.y}
		case "U":
			dest = &coordinate{hPos.x, hPos.y + mag}
		case "D":
			dest = &coordinate{hPos.x, hPos.y - mag}
		}

		for !(*hPos == *dest) {
			hPos.x += utils.Sign(dest.x - hPos.x)
			hPos.y += utils.Sign(dest.y - hPos.y)

			xDist, yDist := hPos.x-tPos.x, hPos.y-tPos.y
			if utils.Abs(xDist) > 1 {
				if utils.Abs(yDist) > 0 {
					tPos.y += yDist
				}
				tPos.x += utils.Sign(xDist)
			} else if utils.Abs(yDist) > 1 {
				if utils.Abs(xDist) > 0 {
					tPos.x += xDist
				}
				tPos.y += utils.Sign(yDist)
			}
			visitedTpos[coordinate{x: tPos.x, y: tPos.y}] = empty
		}
	}
	return strconv.Itoa(len(visitedTpos))
}

func part2() string {

	const k = 10
	dest := &coordinate{x: 0, y: 0}
	knots := [k]*coordinate{}
	for i := 0; i < k; i++ {
		knots[i] = &coordinate{x: 0, y: 0}
	}

	visitedTpos := map[coordinate]struct{}{
		coordinate{knots[k-1].x, knots[k-1].y}: empty,
	}

	for _, move := range input {
		m := strings.Split(move, " ")
		dir := m[0]
		mag, _ := strconv.Atoi(m[1])

		switch dir {
		case "L":
			dest = &coordinate{knots[0].x - mag, knots[0].y}
		case "R":
			dest = &coordinate{knots[0].x + mag, knots[0].y}
		case "U":
			dest = &coordinate{knots[0].x, knots[0].y + mag}
		case "D":
			dest = &coordinate{knots[0].x, knots[0].y - mag}
		}

		for !(*knots[0] == *dest) {
			knots[0].x += utils.Sign(dest.x - knots[0].x)
			knots[0].y += utils.Sign(dest.y - knots[0].y)

			for i := 1; i < k; i++ {
				xDist, yDist := knots[i-1].x-knots[i].x, knots[i-1].y-knots[i].y
				if utils.Abs(xDist) > 1 {
					if utils.Abs(yDist) > 0 {
						knots[i].y += utils.Sign(yDist)
					}
					knots[i].x += utils.Sign(xDist)
				} else if utils.Abs(yDist) > 1 {
					if utils.Abs(xDist) > 0 {
						knots[i].x += utils.Sign(xDist)
					}
					knots[i].y += utils.Sign(yDist)
				} else {
					break
				}
			}
			visitedTpos[coordinate{knots[k-1].x, knots[k-1].y}] = empty
		}
	}
	return strconv.Itoa(len(visitedTpos))
}
