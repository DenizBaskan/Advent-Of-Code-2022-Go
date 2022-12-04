package main

import (
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
	var amount int

    for _, val := range input {
		var firstElfFirstSection, firstElfSecondSection, secondElfFirstSection, secondElfSecondSection int
		_, err := fmt.Sscanf(val, "%d-%d,%d-%d", &firstElfFirstSection, &firstElfSecondSection, &secondElfFirstSection, &secondElfSecondSection)
		if err != nil {
			panic(err)
		}

		if (firstElfFirstSection <= secondElfFirstSection && firstElfSecondSection >= secondElfSecondSection) || (secondElfFirstSection <= firstElfFirstSection && secondElfSecondSection >= firstElfSecondSection) {
			amount++
		}
	}

	return amount
}

func partTwo(input []string) int {
    var amount int

    for _, val := range input {
		var firstElfFirstSection, firstElfSecondSection, secondElfFirstSection, secondElfSecondSection int
		_, err := fmt.Sscanf(val, "%d-%d,%d-%d", &firstElfFirstSection, &firstElfSecondSection, &secondElfFirstSection, &secondElfSecondSection)
		if err != nil {
			panic(err)
		}

		var firstElfRange, secondElfRange []int

		for i := firstElfFirstSection; i < firstElfSecondSection + 1; i++ {
			firstElfRange = append(firstElfRange, i)
		}

		for i := secondElfFirstSection; i < secondElfSecondSection + 1; i++ {
			secondElfRange = append(secondElfRange, i)
		}
		
		Out:
		for _, v1 := range firstElfRange {
			for _, v2 := range secondElfRange {
				if v1 == v2 {
					amount++
					break Out
				}
			}
		}
	}

	return amount
}

func main() {
	var sess aoc.Session

	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}

	input, err := sess.RetrieveInput(2022, 4)
	if err != nil {
		panic(err)
	}
	
	inputLines := input.Strings()
	
	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
