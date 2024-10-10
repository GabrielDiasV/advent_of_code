package main

import (
	"advent_of_code_mod/utils"
	"fmt"
	"strings"
)

const day = 1
const decStartNumb = 48
const decEndNumb = 57

func generateInput() string {
	problemInput := utils.LoadFile(day)
	return problemInput
}

func generateHashmap() map[string]int {
	hashmap := make(map[string]int)
	hashmap["one"] = 1
	hashmap["two"] = 2
	hashmap["three"] = 3
	hashmap["four"] = 4
	hashmap["five"] = 5
	hashmap["six"] = 6
	hashmap["seven"] = 7
	hashmap["eight"] = 8
	hashmap["nine"] = 9
	return hashmap
}

func solvingProblem(problemInput string) int {
	totalValue := 0
	listStrings := strings.Split(problemInput, "\n")
	hashmap := generateHashmap()
	for index := 0; index < len(listStrings); index++ {
		lineString := listStrings[index]
		lastNumbFound := 0
		firstNumbFound := 0
		firstNumbIndex := len(lineString)
		lastNumbIndex := 0
		// checks if a number string is presented on the line
		for numberString, numberValue := range hashmap {
			tempSubsIndex := strings.Index(lineString, numberString)
			if tempSubsIndex <= firstNumbIndex && tempSubsIndex > -1 {
				firstNumbIndex = tempSubsIndex
				firstNumbFound = numberValue
			}
			if tempSubsIndex >= lastNumbIndex && tempSubsIndex > -1 {
				lastNumbIndex = tempSubsIndex
				lastNumbFound = numberValue
			}
		}
		// checks if a int number is presented on the line
		for charIndex, charDec := range lineString {
			// checks if the string is a int number
			if charDec >= decStartNumb && charDec <= decEndNumb {
				if charIndex <= firstNumbIndex {
					firstNumbFound = int(charDec - decStartNumb)
					firstNumbIndex = charIndex
				}
				if charIndex >= lastNumbIndex {
					lastNumbFound = int(charDec - decStartNumb)
					lastNumbIndex = charIndex
				}
			}
		}
		totalValue = totalValue + 10*firstNumbFound
		totalValue = totalValue + lastNumbFound
	}
	return totalValue
}

func main() {
	problemInput := generateInput()
	answer := solvingProblem(problemInput)
	fmt.Print("The answer is: ", answer)
}
