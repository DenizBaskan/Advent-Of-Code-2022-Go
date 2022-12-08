package main

import (
	"strconv"
	"strings"
	"sort"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
	files := make(map[string]*int)
	files["/"] = nil
	currentDir := "/"

	for _, line := range input {
		if line[2] == 'l' && line[3] == 's' {
			continue
		} else if line[2] == 'c' && line[3] == 'd' {
			if line[5] == '.' && line[6] == '.' {
				currentDir = currentDir[:strings.LastIndex(currentDir[:len(currentDir) - 2], "/")] + "/"
			} else if line[5] == '/' {
				currentDir = "/"
			} else {
				currentDir += line[5:] + "/"
			}
		} else if line[0] == 'd' && line[1] == 'i' && line[2] == 'r' {
			files[currentDir + line[4:]] = nil
		} else {
			size, name := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
			intSize, _ := strconv.Atoi(size)
			files[currentDir + name] = &intSize
		}
	}

	var sum int

	for name1, size1 := range files {
		if size1 == nil {
			var sizeOfDir int

			for name2, size2 := range files {
				if len(name2) >= len(name1) && name2[:len(name1)] == name1 && size2 != nil {
					sizeOfDir += *size2
				}
			}

			if sizeOfDir <= 100000 {
				sum += sizeOfDir
			}
		}
	}

	return sum
}

func partTwo(input []string) int {
	files := make(map[string]*int)
	files["/"] = nil
	currentDir := "/"

	for _, line := range input {
		if line[2] == 'l' && line[3] == 's' {
			continue
		} else if line[2] == 'c' && line[3] == 'd' {
			if line[5] == '.' && line[6] == '.' {
				currentDir = currentDir[:strings.LastIndex(currentDir[:len(currentDir) - 2], "/")] + "/"
			} else if line[5] == '/' {
				currentDir = "/"
			} else {
				currentDir += line[5:] + "/"
			}
		} else if line[0] == 'd' && line[1] == 'i' && line[2] == 'r' {
			files[currentDir + line[4:]] = nil
		} else {
			size, name := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
			intSize, _ := strconv.Atoi(size)
			files[currentDir + name] = &intSize
		}
	}

	var total int

	for _, size := range files {
		if size != nil {
			total += *size
		}
	}

	amountToPurge := 30000000 - (70000000 - total)

	var possibleAmounts []int

	for name1, size1 := range files {
		if size1 == nil {
			var sizeOfDir int

			for name2, size2 := range files {
				if len(name2) >= len(name1) && name2[:len(name1)] == name1 && size2 != nil {
					sizeOfDir += *size2
				}
			}

			if sizeOfDir >= amountToPurge {
				possibleAmounts = append(possibleAmounts, sizeOfDir)
			}
		}
	}

	sort.Ints(possibleAmounts)

	return possibleAmounts[0]
}

func main() {
	var sess aoc.Session
	
	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}
	
	input, err := sess.RetrieveInput(2022, 7)
	if err != nil {
		panic(err)
	}
		
	inputLines := input.Strings()
	
	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
