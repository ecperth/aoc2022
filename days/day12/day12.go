package day12

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
)

var input = utils.ReadInputAsBytes(12)

var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

var empty struct{}

var startNode, endNode = findStartAndEnd()

// find start / end and replace with a and z
func findStartAndEnd() (s [2]int, e [2]int) {
	foundStart, foundEnd := false, false
	for !foundStart && !foundEnd {
		for y, line := range input {
			for x, value := range line {
				if value == 'S' {
					s = [2]int{y, x}
					input[y][x] = 'a'
					foundStart = true
				} else if value == 'E' {
					e = [2]int{y, x}
					input[y][x] = 'z'
					foundEnd = true
				}
			}
		}
	}
	return
}

func part1() string {

	//breadth-first search
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maxY, maxX := len(input)-1, len(input[0])-1
	visitedNodes := make(map[[2]int]struct{})
	currentDepth := 0
	nextDepthNodes := map[[2]int]struct{}{}
	currentDepthNodes := map[[2]int]struct{}{startNode: empty}

	for {
		for node, _ := range currentDepthNodes {
			for _, dir := range dirs {
				if node[0]+dir[0] >= 0 && node[0]+dir[0] <= maxY && node[1]+dir[1] >= 0 && node[1]+dir[1] <= maxX {
					if [2]int{node[0] + dir[0], node[1] + dir[1]} == startNode {
						continue
					}
					if input[node[0]+dir[0]][node[1]+dir[1]] <= input[node[0]][node[1]] || input[node[0]+dir[0]][node[1]+dir[1]] == input[node[0]][node[1]]+1 {
						if [2]int{node[0] + dir[0], node[1] + dir[1]} == endNode {
							return strconv.Itoa(currentDepth + 1)
						}
						_, ok := visitedNodes[[2]int{node[0] + dir[0], node[1] + dir[1]}]
						if !ok {
							nextDepthNodes[[2]int{node[0] + dir[0], node[1] + dir[1]}] = empty
						}
					}
				}
			}
			visitedNodes[node] = empty
			delete(currentDepthNodes, node)
		}
		for k, v := range nextDepthNodes {
			currentDepthNodes[k] = v
			delete(nextDepthNodes, k)
		}
		currentDepth++
		if len(currentDepthNodes) == 0 {
			break
		}
	}
	return strconv.Itoa(1)
}

func part2() string {

	//breadth-first search but in reverse
	//start at end and look for first 'a'. This time we can descend a max of 1 tile and go up as many as we want
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maxY, maxX := len(input)-1, len(input[0])-1
	visitedNodes := make(map[[2]int]struct{})
	currentDepth := 0
	nextDepthNodes := map[[2]int]struct{}{}
	currentDepthNodes := map[[2]int]struct{}{endNode: empty}

	for {
		for node, _ := range currentDepthNodes {
			for _, dir := range dirs {
				if node[0]+dir[0] >= 0 && node[0]+dir[0] <= maxY && node[1]+dir[1] >= 0 && node[1]+dir[1] <= maxX {
					if [2]int{node[0] + dir[0], node[1] + dir[1]} == endNode {
						continue
					} else if input[node[0]+dir[0]][node[1]+dir[1]] >= input[node[0]][node[1]] || input[node[0]+dir[0]][node[1]+dir[1]] == input[node[0]][node[1]]-1 {
						if input[node[0]+dir[0]][node[1]+dir[1]] == 'a' {
							return strconv.Itoa(currentDepth + 1)
						}
						_, ok := visitedNodes[[2]int{node[0] + dir[0], node[1] + dir[1]}]
						if !ok {
							nextDepthNodes[[2]int{node[0] + dir[0], node[1] + dir[1]}] = empty
						}
					}
				}
			}
			visitedNodes[node] = empty
			delete(currentDepthNodes, node)
		}
		for k, v := range nextDepthNodes {
			currentDepthNodes[k] = v
			delete(nextDepthNodes, k)
		}
		currentDepth++
		if len(currentDepthNodes) == 0 {
			break
		}
	}
	return strconv.Itoa(1)
}
