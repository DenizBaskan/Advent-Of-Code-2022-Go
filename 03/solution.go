package main

import (
	"strings"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
	var sum int 

    for _, val := range input {
		firstCompartment := val[:len(val) / 2]
		secondCompartment := val[len(val) / 2:]

		var mutualChar rune

		for _, c := range firstCompartment {
			if strings.Contains(secondCompartment, string(c)) {
				mutualChar = c
				break
			}
		}

		if mutualChar > 90 {
			mutualChar = mutualChar - 96
		} else {
			mutualChar = (mutualChar - 64) + 26
		}

		sum += int(mutualChar)
	}

	return sum
}

func partTwo(input []string) int {
    var sum int

    for i := 2; i < len(input); i += 3 {
		var mutualChar rune

		Out:
		for _, v1 := range input[i - 2] {
			for _, v2 := range input[i - 1] {
				for _, v3 := range input[i] {
					if v1 == v2 && v2 == v3 {
						mutualChar = v1
						break Out
					}
				}
			}
		}

		if mutualChar > 90 {
			mutualChar = mutualChar - 96
		} else {
			mutualChar = (mutualChar - 64) + 26
		}

		sum += int(mutualChar)
	}

	return sum
}

func main() {
	var sess aoc.Session

	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}

	input, err := sess.RetrieveInput(2022, 3)
	if err != nil {
		panic(err)
	}
	
	inputLines := input.Strings()

	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
