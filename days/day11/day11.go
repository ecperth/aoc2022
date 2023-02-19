package day11

import (
	"aoc2022/days"
	"aoc2022/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(11)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

const monkeyLen = 7

type test struct {
	divisor int
	passM   int
	failM   int
}

type monkey struct {
	items    []int
	operator rune
	operand  int
	test     test
}

func part1() string {

	rounds := 20
	//read in monkey data
	monkeys := parseMonkeysFromInput(input)

	iCounts := make(map[int]int)
	//begin inspection
	for i := 0; i < rounds; i++ {
		for mi, m := range monkeys {
			for _, item := range m.items {
				var w int
				if m.operator == '*' {
					if m.operand == 0 {
						w = item * item
					} else {
						w = item * m.operand
					}
				} else {
					w = item + m.operand
				}
				w = w / 3

				if w%m.test.divisor == 0 {
					monkeys[m.test.passM].items = append(monkeys[m.test.passM].items, w)
				} else {
					monkeys[m.test.failM].items = append(monkeys[m.test.failM].items, w)

				}
				iCounts[mi] += 1
			}
			monkeys[mi].items = nil
		}
	}
	order := utils.SortMapKeysByValue(iCounts)
	return strconv.Itoa(iCounts[order[len(order)-1]] * iCounts[order[len(order)-2]])
}

func part2() string {

	rounds := 10000
	//read in monkey data
	monkeys := parseMonkeysFromInput(input)

	/*	lowest common multiple
		We need a way of keeping the worry level down. Since all the tests are just checking the (worry level) % (test divisor)
		we should be able to divide by the multiplication of all the divisors to achieve this. This is because of properties:

		(x+z)%y == (x%y)+z
		(x*z)%y == (x%y)*z

		Meaning that the worry level % a monkeys divisor would not be impacted if we divided the worry level by the monkeys divisor
		before or after each turn. Extending that to all monkeys => if we divide the worry level by the divisors of all the monkeys
		before or after each turn then the order will be preserved.
	*/
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test.divisor
	}

	iCounts := make(map[int]int)
	//begin inspection
	for i := 0; i < rounds; i++ {
		for mi, m := range monkeys {
			for _, item := range m.items {
				var w int
				if m.operator == '*' {
					if m.operand == 0 {
						w = item * item
					} else {
						w = item * m.operand
					}
				} else {
					w = item + m.operand
				}
				w %= lcm

				if w%m.test.divisor == 0 {
					monkeys[m.test.passM].items = append(monkeys[m.test.passM].items, w)
				} else {
					monkeys[m.test.failM].items = append(monkeys[m.test.failM].items, w)

				}
				iCounts[mi] += 1
			}
			monkeys[mi].items = nil
		}
	}
	order := utils.SortMapKeysByValue(iCounts)
	return strconv.Itoa(iCounts[order[len(order)-1]] * iCounts[order[len(order)-2]])
}

func parseMonkeysFromInput(input []string) []monkey {

	monkeyCount := (len(input) + 1) / monkeyLen
	monkeys := make([]monkey, monkeyCount, monkeyCount)
	for i := 0; i < monkeyCount; i++ {
		m := monkey{
			items: []int{},
		}

		//items
		for _, startingItems := range strings.Split(strings.Split(input[i*monkeyLen+1], ":")[1], ",") {
			st, _ := strconv.Atoi(strings.Trim(startingItems, " "))
			if st != 0 {
				m.items = append(m.items, st)
			}
		}

		//operation
		tokens := strings.Split(strings.Trim(strings.Split(input[i*monkeyLen+2], "=")[1], " "), " ")
		m.operator = rune(tokens[1][0])
		operand, err := strconv.Atoi(tokens[2])
		if err == nil {
			m.operand = operand
		}

		//tests
		var d, p, f int
		utils.SscanfUnsafe(strings.TrimSpace(input[i*monkeyLen+3]), "Test: divisible by %d", &d)
		utils.SscanfUnsafe(strings.TrimSpace(input[i*monkeyLen+4]), "If true: throw to monkey %d", &p)
		utils.SscanfUnsafe(strings.TrimSpace(input[i*monkeyLen+5]), "If false: throw to monkey %d", &f)
		m.test = test{d, p, f}

		monkeys[i] = m
	}

	return monkeys
}
