package day7

import (
	"aoc2022/days"
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(7)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

const (
	CD    = "cd"
	DIR   = "dir"
	INPUT = "$"
)

type dir struct {
	name    string
	subDirs map[string]*dir
	files   map[string]int
}

func part1() string {

	//Build dir structure
	fileSystem := buildFileSystemFromInput(input)

	result := 0
	subDirsAndSizes := getSubdirsAndSizes(fileSystem.subDirs["/"])
	for _, v := range subDirsAndSizes {
		if v <= 100000 {
			result += v
		}
	}

	return strconv.Itoa(result)
}

func part2() string {

	//Build dir structure
	fileSystem := buildFileSystemFromInput(input)

	subDirsAndSizes := getSubdirsAndSizes(fileSystem.subDirs["/"])
	totalDirSizes := subDirsAndSizes[fileSystem.subDirs["/"]]
	freeSpace := 70000000 - totalDirSizes

	subDirSizes := make([]int, 0, len(subDirsAndSizes))
	for _, v := range subDirsAndSizes {
		subDirSizes = append(subDirSizes, v)
	}
	sort.Ints(subDirSizes)

	for _, subDirSize := range subDirSizes {
		if subDirSize > 30000000-freeSpace {
			return strconv.Itoa(subDirSize)
		}
	}

	return "Not found"
}

func buildFileSystemFromInput(input []string) (fileSystem *dir) {
	wd := fileSystem
	wd.subDirs = map[string]*dir{"/": {"/", map[string]*dir{}, map[string]int{}}}

	for _, line := range input {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case INPUT:
			if tokens[1] == CD {
				dirName := tokens[2]
				dest, ok := wd.subDirs[dirName]
				if !ok {
					panic(fmt.Errorf("dir %s does not exist in current dir", dirName))
				} else {
					wd = dest
				}
			}
		case DIR:
			dirName := tokens[1]
			_, ok := wd.subDirs[dirName]
			if !ok {
				wd.subDirs[dirName] = &dir{dirName, map[string]*dir{"..": wd}, map[string]int{}}
			}
		default:
			fileName := tokens[1]
			_, ok := wd.files[fileName]
			if !ok {
				fileSize, _ := strconv.Atoi(tokens[0])
				wd.files[fileName] = fileSize
			}
		}
	}
	return fileSystem
}

/*
returns a map where the key is the pointer to a dir and the value is the total size of the dir
*/
func getSubdirsAndSizes(currentDir *dir) (result map[*dir]int) {
	dirSize := 0

	for name, subDir := range currentDir.subDirs {
		if name == ".." {
			continue
		}
		subDirSizes := getSubdirsAndSizes(subDir)
		for k, v := range subDirSizes {
			result[k] = v
			if k == subDir {
				dirSize += v
			}
		}
	}

	for _, fileSize := range currentDir.files {
		dirSize += fileSize
	}
	result[currentDir] = dirSize
	return result
}
