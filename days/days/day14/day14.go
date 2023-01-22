package day14

import (
	"aoc2022/days"
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var input = utils.ReadInputAsStrings(14)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

var empty struct{}

const displayHeight = 55

var offset = 0
var screen []byte

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

	rocks, maxY, minX, maxX := makeRocksMap(input)
	sand := make(map[[2]int]struct{})
	initialiseScreen(maxY, minX, maxX, rocks)
	reachedAbyss := false
	for !reachedAbyss {
		currentSandPos := [2]int{500, 0}
		for {
			//if the y pos of falling sand is = maxY we are done
			if currentSandPos[1] >= maxY {
				reachedAbyss = true
				break
			}
			nextSandPos, moved := moveSand(currentSandPos, rocks, sand)
			if !moved {
				sand[currentSandPos] = empty
				break
			}
			updateScreen(minX, maxX, currentSandPos, nextSandPos)
			currentSandPos = nextSandPos
		}
		drawScreen(minX, maxX, len(sand))
	}
	return strconv.Itoa(len(sand))
}

func part2() string {

	rocks, maxY, _, _ := makeRocksMap(input)
	sand := make(map[[2]int]struct{})
	for {
		startPoint := [2]int{500, 0}
		//if the array of at rest sand contains 500,0 we are done
		if utils.Contains(sand, startPoint) {
			break
		}
		currentSandPos := startPoint
		for {
			//if sand y pos is 1 greater than the maxY then stop falling and move to next grain
			if currentSandPos[1] >= maxY+1 {
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
func makeRocksMap(input []string) (map[[2]int]struct{}, int, int, int) {
	rocks := make(map[[2]int]struct{})
	maxY, minX, maxX := 0, 1000, 0
	for _, row := range input {
		ls := strings.Split(row, " -> ")

		left, right, _ := strings.Cut(ls[0], ",")
		x1, y1 := utils.AtoiUnsafe(left), utils.AtoiUnsafe(right)
		maxY, minX, maxX = utils.Max(maxY, y1), utils.Min(minX, x1), utils.Max(maxX, x1)

		for i := 1; i < len(ls); i++ {
			left, right, _ = strings.Cut(ls[i], ",")
			x2, y2 := utils.AtoiUnsafe(left), utils.AtoiUnsafe(right)
			for _, point := range interpolate([2]int{x1, y1}, [2]int{x2, y2}) {
				rocks[point] = empty
			}
			x1, y1 = x2, y2
			maxY, minX, maxX = utils.Max(maxY, y1), utils.Min(minX, x1), utils.Max(maxX, x1)
		}
	}
	return rocks, maxY, minX, maxX
}

// receives the current position of falling grain of sand, the rocks, already fallen sand.
// returns next position of passed in grain of sand and if the sand moved or not.
func moveSand(currentSandPos [2]int, rocks, sand map[[2]int]struct{}) ([2]int, bool) {
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

func initialiseScreen(maxY, minX, maxX int, rocks map[[2]int]struct{}) {
	width := maxX - minX + 1
	screen = make([]byte, (width+1)*(maxY+1))
	for y := 0; y <= maxY; y++ {

		for x := minX; x <= maxX; x++ {
			pixel := [2]int{x, y}
			c := byte('.')
			if utils.Contains(rocks, pixel) {
				c = '|'
			}
			screen[y*(width+1)+x-minX] = c
		}
		screen[(y+1)*(width+1)-1] = 10
	}
	screen[len(screen)-1] = 10
}

func updateScreen(minX, maxX int, lastSandPos, currentSandPos [2]int) {
	width := maxX - minX + 1
	if currentSandPos[0] >= minX && currentSandPos[0] <= maxX {
		screen[lastSandPos[1]*(width+1)+lastSandPos[0]-minX] = '.'
		screen[currentSandPos[1]*(width+1)+currentSandPos[0]-minX] = 'o'
	}
	if currentSandPos[1] >= (offset+1)*displayHeight {
		offset++
	}
}

func drawScreen(minX, maxX int, sandCount int) {

	width := maxX - minX + 1
	startIndex := (width + 1) * displayHeight * offset
	endIndex := startIndex + (width+1)*displayHeight
	if endIndex >= len(screen) {
		endIndex = len(screen)
	}

	utils.ClearTerminal()
	fmt.Print(string(screen[startIndex:endIndex]))
	fmt.Printf("sand count: %d\n", sandCount)
	time.Sleep(50 * time.Millisecond)
}
