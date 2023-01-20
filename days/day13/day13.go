package day13

import (
	"aoc2022/days"
	"aoc2022/utils"
	"encoding/json"
	"sort"
	"strconv"
)

var input = utils.ReadInputAsStrings(13)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	parsedPackets := parsePackets(input)
	result := 0
	for i := 0; i < len(parsedPackets); i += 2 {
		if compareParts(parsedPackets[i], parsedPackets[i+1]) >= 0 {
			result += i/2 + 1
		}
	}
	return strconv.Itoa(result)
}

func part2() string {

	result := 1
	parsedPackets := parsePackets(input)

	decoderKey1, decoderKey2 := []interface{}{float64(2)}, []interface{}{float64(6)}
	parsedPackets = append(parsedPackets, decoderKey1)
	parsedPackets = append(parsedPackets, decoderKey2)

	sort.SliceStable(parsedPackets, func(i, j int) bool {
		return compareParts(parsedPackets[i], parsedPackets[j]) >= 0
	})

	for i, packet := range parsedPackets {
		if compareParts(packet, decoderKey1) == 0 || compareParts(packet, decoderKey2) == 0 {
			result *= i + 1
		}
	}

	return strconv.Itoa(result)
}

func parsePackets(packets []string) (parsedPackets [][]interface{}) {
	for _, packet := range packets {
		if len(packet) > 0 {
			var parsedPacket []interface{}
			json.Unmarshal([]byte(packet), &parsedPacket)
			parsedPackets = append(parsedPackets, parsedPacket)
		}
	}
	return
}

func compareParts(left, right []interface{}) int {
	for l := 0; l < len(left); l++ {

		//for case everything is equal to this point but right runs out of parts
		if l > len(right)-1 {
			return -1
		}

		leftNum, isLeftNum := left[l].(float64)
		rightNum, isRightNum := right[l].(float64)

		leftList, isLeftList := left[l].([]interface{})
		rightList, isRightList := right[l].([]interface{})
		c := 0
		if isLeftNum && isRightNum {
			c = utils.Sign(rightNum - leftNum)
		} else if isLeftList && isRightList {
			//both are slices and equal then continue
			c = compareParts(leftList, rightList)
		} else {
			if isLeftNum {
				leftList = []interface{}{leftNum}
			} else if isRightNum {
				rightList = []interface{}{rightNum}
			}
			c = compareParts(leftList, rightList)
		}
		//if current part is equal then continue
		if c != 0 {
			return c
		}
	}
	//for case everything is equal to this point. Will return 0 if both are exactly the same and 1 if right was longer
	return utils.Sign(len(right) - len(left))
}
