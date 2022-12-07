package main

import (
	"math/rand"
	"strconv"
	"strings"
	"sort"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

type File struct {
	Size       int
	ParentUUID int64
	UUID       int64     
	Contents   map[string]File
}

func recursivelyGetTotal(dir File, sum *int) {
	for _, f := range dir.Contents {
		if f.UUID != 0 {
			recursivelyGetTotal(f, sum)
		} else {
			*sum += f.Size
		}
	}
}

func partOne(input []string) int {
	currentDir := File{Size: 0, ParentUUID: 0, UUID: 99999999 - rand.Int63n(90000000), Contents: make(map[string]File)}
	rootDir := &currentDir

	dirs := []File{currentDir}
	
	for _, line := range input {
		if line[2] == 'l' && line[3] == 's' {
			continue
		} else if line[2] == 'c' && line[3] == 'd' {
			if line[5] == '.' && line[6] == '.' {
				for _, dir := range dirs {
					if dir.UUID == currentDir.ParentUUID {
						currentDir = dir
					}
				}
			} else if line[5] == '/' {
				currentDir = *rootDir
			} else {
				currentDir = currentDir.Contents[line[5:]]
			}
		} else if line[0] == 'd' && line[1] == 'i' && line[2] == 'r' {
			newDir := File{Size: 0, ParentUUID: currentDir.UUID, UUID: 99999999 - rand.Int63n(90000000), Contents: make(map[string]File)}
			currentDir.Contents[line[4:]] = newDir
			dirs = append(dirs, newDir)
		} else {
			size, name := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
			intSize, err := strconv.Atoi(size)
			if err != nil {
				panic(err)
			}
			currentDir.Contents[name] = File{Size: intSize, ParentUUID: 0, UUID: 0, Contents: nil}
		}
	}

	var sum int

	for _, dir := range dirs {
		var dirTotal int
		recursivelyGetTotal(dir, &dirTotal)
		if dirTotal <= 100000 {
			sum += dirTotal
		}
	}

	return sum
}

func partTwo(input []string) int {
	currentDir := File{Size: 0, ParentUUID: 0, UUID: 99999999 - rand.Int63n(90000000), Contents: make(map[string]File)}
	rootDir := &currentDir

	dirs := []File{currentDir}
	
	for _, line := range input {
		if line[2] == 'l' && line[3] == 's' {
			continue
		} else if line[2] == 'c' && line[3] == 'd' {
			if line[5] == '.' && line[6] == '.' {
				for _, dir := range dirs {
					if dir.UUID == currentDir.ParentUUID {
						currentDir = dir
					}
				}
			} else if line[5] == '/' {
				currentDir = *rootDir
			} else {
				currentDir = currentDir.Contents[line[5:]]
			}
		} else if line[0] == 'd' && line[1] == 'i' && line[2] == 'r' {
			newDir := File{Size: 0, ParentUUID: currentDir.UUID, UUID: 99999999 - rand.Int63n(90000000), Contents: make(map[string]File)}
			currentDir.Contents[line[4:]] = newDir
			dirs = append(dirs, newDir)
		} else {
			size, name := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
			intSize, err := strconv.Atoi(size)
			if err != nil {
				panic(err)
			}
			currentDir.Contents[name] = File{Size: intSize, ParentUUID: 0, UUID: 0, Contents: nil}
		}
	}

	var amounts []int
	var rootTotal int

	for _, dir := range dirs {
		var dirTotal int
		recursivelyGetTotal(dir, &dirTotal)
		amounts = append(amounts, dirTotal)
		if dirTotal > rootTotal {
			rootTotal = dirTotal
		}
	}

	var amountToFree = 30000000 - (70000000 - rootTotal)

	var possibleAmounts []int 

	for _, amount := range amounts {
		if amount >= amountToFree {
			possibleAmounts = append(possibleAmounts, amount)
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
