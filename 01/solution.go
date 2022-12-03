package main

import (
	"strconv"
	"sort"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
	var highestScore, buffer int

	for _, val := range input {
		if val == "" {
			if buffer > highestScore {
				highestScore = buffer
			}
			buffer = 0
		} else {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			buffer += intVal
		}
	}

	return highestScore
}

func partTwo(input []string) int {
	var buffer int
	var scores []int

	for _, val := range input {
		if val == "" {
			scores = append(scores, buffer)
			buffer = 0
		} else {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			buffer += intVal
		}
	}

	sort.Ints(scores)

	return scores[len(scores) - 1] + scores[len(scores) - 2] + scores[len(scores) - 3]
}

func main() {
	var sess aoc.Session

	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}

	input, err := sess.RetrieveInput(2022, 1)
	if err != nil {
		panic(err)
	}
	
	inputLines := input.Strings()

	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
