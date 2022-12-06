package main

import (
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input string) int {
	var i int

  	Out:
	for i = 3; i < len(input); i++ {
		var previousFourChars string

		for j := i - 3; j != i + 1; j++ {
			previousFourChars += string(input[j])
		}

		for i1, c1 := range previousFourChars {
			for i2, c2 := range previousFourChars {
				if c1 == c2 && i1 != i2 {
					continue Out
				}
			}
		}

		break
	}

	return i + 1
}

func partTwo(input string) int {
	var i int

	Out:
	for i = 13; i < len(input); i++ {
		var previousFourChars string

		for j := i - 13; j != i + 1; j++ {
			previousFourChars += string(input[j])
		}

		for i1, c1 := range previousFourChars {
			for i2, c2 := range previousFourChars {
				if c1 == c2 && i1 != i2 {
					continue Out
				}
			}
		}

		break
	}

	return i + 1
}

func main() {
	var sess aoc.Session
	
	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}
	
	input, err := sess.RetrieveInput(2022, 6)
	if err != nil {
		panic(err)
	}
		
	inputRaw := input.Raw()

	fmt.Printf("Part one: %d\n", partOne(inputRaw))
	fmt.Printf("Part two: %d\n", partTwo(inputRaw))
}
