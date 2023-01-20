package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func ReadInputAsStrings(day int) []string {
	s := getScanner(day)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func ReadInputAsBytes(day int) [][]byte {
	s := getScanner(day)

	var matrix [][]byte
	for s.Scan() {
		matrix = append(matrix, []byte(s.Text()))
	}
	return matrix
}

func getScanner(day int) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("./inputs/day%d", day))
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign[T Number](x T) int {
	if x == 0 {
		return 0
	} else if x > 0 {
		return 1
	} else {
		return -1
	}
}

func SortMapKeysByValue[kT comparable, vT int](m map[kT]vT) []kT {
	keys := make([]kT, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
