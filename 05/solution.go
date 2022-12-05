package main

import (
	"unicode"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) string {
	var stacks = make(map[int][]string)

	var stackTableHeight, stackTableWidth int 

	for ; stackTableHeight < len(input); stackTableHeight++ {
		if input[stackTableHeight][1] == '1' {
			break
		}
	}

	for _, v := range input[stackTableHeight] {
		if unicode.IsDigit(v) {
			stackTableWidth++
		}
	}

	for i := 1; i < stackTableWidth + 1; i++ {
		var stack []string

		for j := stackTableHeight - 1; j != -1; j-- {
			v := string(input[j][(3 * (i - 1)) + i])
			if v != " " {
				stack = append(stack, v)
			}
		}

		stacks[i] = stack
	}

	for i := stackTableHeight + 2; i < len(input); i++ {
		var amountToMove, originalStack, targetStack int 
		_, err := fmt.Sscanf(input[i], "move %d from %d to %d", &amountToMove, &originalStack, &targetStack)
		if err != nil {
			panic(err)
		}

		var toBeRemoved []string 

		for j := 0; j < amountToMove; j++ {
			toBeRemoved = append(toBeRemoved, stacks[originalStack][len(stacks[originalStack]) - 1 - j])
		}

		stacks[originalStack] = stacks[originalStack][:len(stacks[originalStack]) - amountToMove]
		stacks[targetStack] = append(stacks[targetStack], toBeRemoved...)
	}

	var topOfEachStackString string 

	for i := 1; i < stackTableWidth + 1; i++ {
		topOfEachStackString += stacks[i][len(stacks[i]) - 1]
	}

	return topOfEachStackString
}

func partTwo(input []string) string {
	var stacks = make(map[int][]string)

	var stackTableHeight, stackTableWidth int 

	for ; stackTableHeight < len(input); stackTableHeight++ {
		if input[stackTableHeight][1] == '1' {
			break
		}
	}

	for _, v := range input[stackTableHeight] {
		if unicode.IsDigit(v) {
			stackTableWidth++
		}
	}

	for i := 1; i < stackTableWidth + 1; i++ {
		var stack []string

		for j := stackTableHeight - 1; j != -1; j-- {
			v := string(input[j][(3 * (i - 1)) + i])
			if v != " " {
				stack = append(stack, v)
			}
		}

		stacks[i] = stack
	}

	for i := stackTableHeight + 2; i < len(input); i++ {
		var amountToMove, originalStack, targetStack int 
		_, err := fmt.Sscanf(input[i], "move %d from %d to %d", &amountToMove, &originalStack, &targetStack)
		if err != nil {
			panic(err)
		}

		var toBeRemoved []string

		for j := amountToMove - 1; j != -1; j-- {
			toBeRemoved = append(toBeRemoved, stacks[originalStack][len(stacks[originalStack]) - 1 - j])
		}

		stacks[originalStack] = stacks[originalStack][:len(stacks[originalStack]) - amountToMove]
		stacks[targetStack] = append(stacks[targetStack], toBeRemoved...)
	}

	var topOfEachStackString string

	for i := 1; i < stackTableWidth + 1; i++ {
		topOfEachStackString += stacks[i][len(stacks[i]) - 1]
	}

	return topOfEachStackString
}

func main() {
	var sess aoc.Session
	
	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}
	
	input, err := sess.RetrieveInput(2022, 5)
	if err != nil {
		panic(err)
	}
		
	inputLines := input.Strings()

	fmt.Printf("Part one: %s\n", partOne(inputLines))
	fmt.Printf("Part two: %s\n", partTwo(inputLines))
}
