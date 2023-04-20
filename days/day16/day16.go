package day16

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(16)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

type Valve struct {
	flowRate     int
	linkedValves []string
}

type NodeDetails struct {
	pressureReleased int
	timeRemaining    int
}

func part1() string {

	valveMap := make(map[string]Valve)
	var currentValve string
	for _, line := range input {
		var valve, linkedValves string
		var flowRate int
		utils.SscanfUnsafe(line, "Valve %s has flow rate=%d", &valve, &flowRate)
		if currentValve == "" {
			currentValve = valve
		}
		linkedValves = strings.SplitN(strings.Split(line, "valve")[1], " ", 2)[1]
		valveMap[valve] = Valve{flowRate: flowRate, linkedValves: strings.Split(linkedValves, ", ")}
	}

	timeRemaining := 30
	NodeDetailsMap := make(map[string]NodeDetails)
	NodeDetailsMap[currentValve] = NodeDetails{0, timeRemaining}

	for _, connectedValve := range valveMap[currentValve].linkedValves {
		currentMaxPressure, ok := NodeDetailsMap[connectedValve]
		if !ok {
			NodeDetailsMap[connectedValve] = NodeDetails{
				NodeDetailsMap[connectedValve].pressureReleased + timeRemaining*valveMap[connectedValve].flowRate,
				timeRemaining - 2,
			}
		}
	}

	return strconv.Itoa(1)
}

func part2() string {
	return strconv.Itoa(1)
}
