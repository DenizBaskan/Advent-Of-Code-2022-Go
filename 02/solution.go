package main

import (
	"strings"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
    var score int 
    
    for _, val := range input {
        first := strings.Split(val, " ")[0]
        second := strings.Split(val, " ")[1]
        
        if first == "A" && second == "Y" {
            score += 8
        } else if first == "A" && second == "X" {
            score += 4
        } else if first == "A" && second == "Z" {
            score += 3
        } else if first == "B" && second == "Y" {
            score += 5
        } else if first == "B" && second == "X" {
            score++
        } else if first == "B" && second == "Z" {
            score += 9
        } else if first == "C" && second == "Y" {
            score += 2
        } else if first == "C" && second == "X" {
            score += 7
        } else if first == "C" && second == "Z" {
            score += 6
        }
    }
    
    return score
}

func partTwo(input []string) int {
    var score int 
    
    for _, val := range input {
        first := strings.Split(val, " ")[0]
        second := strings.Split(val, " ")[1]
        
        if second == "X" {
            if first == "A" {
                score += 3
            } else if first == "B" {
                score++
            } else if first == "C" {
                score += 2
            }
        } else if second == "Y" {
            score += 3
            if first == "A" {
                score++
            } else if first == "B" {
                score += 2
            } else if first == "C" {
                score += 3
            }
        } else if second == "Z" {
            score += 6
            if first == "A" {
                score += 2
            } else if first == "B" {
                score += 3
            } else if first == "C" {
                score++
            }
        }
    }
    
    return score
}

func main() {
	var sess aoc.Session

	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}

	input, err := sess.RetrieveInput(2022, 2)
	if err != nil {
		panic(err)
	}
	
	inputLines := input.Strings()

	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
