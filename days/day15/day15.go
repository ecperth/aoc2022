package day15

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
)

var input = utils.ReadInputAsStrings(15)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	const yVal = 2000000
	sensorBeaconDistMap := make(map[[2][2]int]int)
	var minX, minY int
	utils.SscanfUnsafe(input[0], "Sensor at x=%d, y=%d", &minX, &minY)
	maxX, maxY := minX, minY

	//Read in sensor-beacon data
	for _, line := range input {
		var sX, sY, bX, bY int
		utils.SscanfUnsafe(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)
		sensorBeaconPair := [2][2]int{{sX, sY}, {bX, bY}}
		dist := utils.Abs(sX-bX) + utils.Abs(sY-bY)
		minX, maxX, minY, maxY = utils.Min(minX, sX-dist), utils.Max(maxX, sX+dist), utils.Min(minY, sY-dist), utils.Max(maxY, sY+dist)
		sensorBeaconDistMap[sensorBeaconPair] = utils.Abs(sX-bX) + utils.Abs(sY-bY)
	}

	result := 0
	//For each x on line y=yVal
	for x := minX; x <= maxX; x++ {
		//For each beacon
		for sensorBeaconPair, dist := range sensorBeaconDistMap {
			sensor, beacon := sensorBeaconPair[0], sensorBeaconPair[1]
			//There is a beacon here
			if beacon[0] == x && beacon[1] == yVal {
				break
			}
			d1 := utils.Abs(sensor[0]-x) + utils.Abs(sensor[1]-yVal)
			//There can not be beacon here as it is with range of sensor
			if d1 <= dist {
				result++
				break
			}
		}
	}
	return strconv.Itoa(result)
}

func part2() string {

	const MaxVal = 4000000
	sensorDistMap := make(map[[2]int]int)
	var distressBeacon [2]int

	//Read in sensor-beacon data
	for _, line := range input {
		var sX, sY, bX, bY int
		utils.SscanfUnsafe(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)
		sensor := [2]int{sX, sY}
		sensorDistMap[sensor] = utils.Abs(sX-bX) + utils.Abs(sY-bY)
	}

	//For each sensor
	for s1, d1 := range sensorDistMap {
		//For each x within its range
		for dx := -(d1 + 1); dx <= (d1 + 1); dx++ {
			cantBe := false
			x := s1[0] + dx
			if x < 0 || x > MaxVal {
				continue
			}
			//For the 2 possible corresponding y values
			for _, y := range []int{s1[1] + (d1 + 1 - utils.Abs(dx)), s1[1] - (d1 + 1 - utils.Abs(dx))} {
				if y > MaxVal || y < 0 {
					continue
				}
				//For each other sensor
				for s2, d2 := range sensorDistMap {
					//If the point with the current x y is within sensor2s range
					if utils.Abs(s2[0]-x)+utils.Abs(s2[1]-y) <= d2 {
						cantBe = true
						break
					}
				}
				if !cantBe {
					distressBeacon = [2]int{x, y}
					break
				}
			}
		}
	}
	return strconv.Itoa(distressBeacon[0]*4000000 + distressBeacon[1])
}
