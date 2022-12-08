package main

import (
	"strconv"
	"sort"
	"fmt"

	aoc "github.com/DenizBaskan/Aoc-Helper"
)

func partOne(input []string) int {
	var amountVisible int

	for i1 := range input {
		for i2 := range input[i1] {
			if i1 == 0 || i1 == len(input) - 1 || i2 == len(input[i1]) - 1 || i2 == 0 {
				amountVisible++
				continue
			}

			if i1 != 0 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToTop int
				for i := 1; i < i1 + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1 - i][i2]))
					if tree > intVal {
						amountVisibleToTop++
					}
				}

				if amountVisibleToTop == i1 {
					amountVisible++
					continue
				}
			}

			if i1 != len(input)-1 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToBottom int
				offsetToBottom := (len(input) - 1) - i1

				for i := 1; i < offsetToBottom + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1 + i][i2]))

					if tree > intVal {
						amountVisibleToBottom++
					}
				}

				if amountVisibleToBottom == offsetToBottom {
					amountVisible++
					continue
				}
			}

			if i2 != 0 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToLeftEdge int

				for i := 1; i < i2 + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1][i2 - i]))

					if tree > intVal {
						amountVisibleToLeftEdge++
					}
				}

				if amountVisibleToLeftEdge == i2 {
					amountVisible++
					continue
				}
			}

			if i2 != len(input[i1]) - 1 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToRightEdge int
				offsetToRight := (len(input) - 1) - i2

				for i := 1; i < offsetToRight + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1][i2 + i]))

					if tree > intVal {
						amountVisibleToRightEdge++
					}
				}

				if amountVisibleToRightEdge == offsetToRight {
					amountVisible++
					continue
				}
			}
		}
	}

	return amountVisible
}

func partTwo(input []string) int {
	var scenicScores []int

	for i1 := range input {
		for i2 := range input[i1] {
			var scores []int 

			if i1 != 0 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToTop int
				for i := 1; i < i1 + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1 - i][i2]))

					amountVisibleToTop++

					if !(tree > intVal) {
						break
					}
				}
				scores = append(scores, amountVisibleToTop)
			} else {
				scores = append(scores, 0)
			}

			if i1 != len(input) - 1 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToBottom int
				offsetToBottom := (len(input) - 1) - i1

				for i := 1; i < offsetToBottom + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1 + i][i2]))

					amountVisibleToBottom++

					if !(tree > intVal) {
						break
					}
				}
				scores = append(scores, amountVisibleToBottom)
			} else {
				scores = append(scores, 0)
			}

			if i2 != 0 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToLeftEdge int

				for i := 1; i < i2 + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1][i2 - i]))

					amountVisibleToLeftEdge++
					
					if !(tree > intVal) {
						break
					}
				}
				scores = append(scores, amountVisibleToLeftEdge)
			} else {
				scores = append(scores, 0)
			}

			if i2 != len(input[i1]) - 1 {
				tree, _ := strconv.Atoi(string(input[i1][i2]))

				var amountVisibleToRightEdge int
				offsetToRight := (len(input) - 1) - i2

				for i := 1; i < offsetToRight + 1; i++ {
					intVal, _ := strconv.Atoi(string(input[i1][i2 + i]))

					amountVisibleToRightEdge++
					
					if !(tree > intVal) {
						break
					}
				}
				scores = append(scores, amountVisibleToRightEdge)
			} else {
				scores = append(scores, 0)
			}

			scenicScores = append(scenicScores, scores[0] * scores[1] * scores[2] * scores[3])
		}
	}

	sort.Ints(scenicScores)

	return scenicScores[len(scenicScores) - 1]
}

func main() {
	var sess aoc.Session
	
	err := sess.InitSessionFromFile("../session.txt")
	if err != nil {
		panic(err)
	}
	
	input, err := sess.RetrieveInput(2022, 8)
	if err != nil {
		panic(err)
	}
		
	inputLines := input.Strings()

	fmt.Printf("Part one: %d\n", partOne(inputLines))
	fmt.Printf("Part two: %d\n", partTwo(inputLines))
}
