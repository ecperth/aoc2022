package day14

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(14)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

var empty struct{}

func interpolate(p1, p2 [2]int) (points [][2]int) {
	n := 0
	if p1[0] == p2[0] {
		dir := utils.Sign(p2[1] - p1[1])
		for {
			nextY := p1[1] + dir*n
			points = append(points, [2]int{p1[0], nextY})
			if nextY == p2[1] {
				break
			}
			n++
		}
	} else {
		dir := utils.Sign(p2[0] - p1[0])
		for {
			nextX := p1[0] + dir*n
			points = append(points, [2]int{nextX, p1[1]})
			if nextX == p2[0] {
				break
			}
			n++
		}
	}
	return points
}

func part1() string {

	rocks, maxY := makeRocksMap(input)
	sand := make(map[[2]int]struct{})
	reachedAbyss := false
	for !reachedAbyss {
		currentSandPos := [2]int{500, 0}
		for {
			//if the y pos of falling sand is = maxY we are done
			if currentSandPos[1] == maxY {
				reachedAbyss = true
				break
			}
			nextSandPos, moved := moveSand(currentSandPos, rocks, sand)
			if !moved {
				sand[currentSandPos] = empty
				break
			}
			currentSandPos = nextSandPos
		}
	}
	return strconv.Itoa(len(sand))
}

func part2() string {

	rocks, maxY := makeRocksMap(input)
	sand := make(map[[2]int]struct{})
	for {
		startPoint := [2]int{500, 0}
		//if the array of at rest sand contains 500,0 we are done
		if utils.Contains(sand, startPoint) {
			break
		}
		currentSandPos := startPoint
		for {
			//if sand y pos is i passed the maxY then stop falling and move to next grain
			if currentSandPos[1] == maxY+1 {
				sand[currentSandPos] = empty
				break
			}
			nextSandPos, moved := moveSand(currentSandPos, rocks, sand)
			if !moved {
				sand[currentSandPos] = empty
				break
			}
			currentSandPos = nextSandPos
		}
	}
	return strconv.Itoa(len(sand))
}

// create a set of xy coords which represent the rocks in the cavern
func makeRocksMap(input []string) (map[[2]int]struct{}, int) {
	rocks := make(map[[2]int]struct{})
	maxY := 0
	for _, row := range input {
		ls := strings.Split(row, " -> ")

		left, right, _ := strings.Cut(ls[0], ",")
		x1, y1 := utils.AtoiUnsafe(left), utils.AtoiUnsafe(right)
		maxY = utils.Max(maxY, y1)

		for i := 1; i < len(ls); i++ {
			left, right, _ = strings.Cut(ls[i], ",")
			x2, y2 := utils.AtoiUnsafe(left), utils.AtoiUnsafe(right)
			for _, point := range interpolate([2]int{x1, y1}, [2]int{x2, y2}) {
				rocks[point] = empty
			}
			x1, y1 = x2, y2
			maxY = utils.Max(maxY, y1)
		}
	}
	return rocks, maxY
}

// receives the current position of falling grain of sand, the rocks, already fallen sand.
// returns next position of passed in grain of sand and if the sand moved or not.
func moveSand(currentSandPos [2]int, rocks map[[2]int]struct{}, sand map[[2]int]struct{}) ([2]int, bool) {
	//down
	nextSandPos := [2]int{currentSandPos[0], currentSandPos[1] + 1}
	if !utils.Contains(rocks, nextSandPos) && !utils.Contains(sand, nextSandPos) {
		return nextSandPos, true
	}
	//left diagonal
	nextSandPos = [2]int{currentSandPos[0] - 1, currentSandPos[1] + 1}
	if !utils.Contains(rocks, nextSandPos) && !utils.Contains(sand, nextSandPos) {
		return nextSandPos, true
	}
	//right diagonal
	nextSandPos = [2]int{currentSandPos[0] + 1, currentSandPos[1] + 1}
	if !utils.Contains(rocks, nextSandPos) && !utils.Contains(sand, nextSandPos) {
		return nextSandPos, true
	}
	//could not move
	return currentSandPos, false
}
