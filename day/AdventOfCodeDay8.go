package day

import (
	"fmt"
	"strconv"
	"strings"
)

func getLeftOrRightInstruction(leftRightInstructions string, leftRightInstructionCounter *int) string {
	returnVal := string(leftRightInstructions[(*leftRightInstructionCounter)])

	if len(leftRightInstructions)-1 == *leftRightInstructionCounter {
		*leftRightInstructionCounter = 0
	} else {
		*leftRightInstructionCounter++
	}

	return returnVal
}

func ProcessDay8(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	if aocDay.Part == "A" {
		processDay8A(inputLines)
	} else {
		processDay8B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay8A(inputLines []string) {
	fmt.Println("*** Start part A ***")
	leftRightInstructions := ""
	directions := make([]string, 0)

	// Read in the input file and split out the important pieces
	for i := 0; i < len(inputLines); i++ {

		inputLine := inputLines[i]
		if len(inputLine) == 0 {
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "=") {
			directions = append(directions, inputLine)
			continue
		} else {
			leftRightInstructions = inputLines[i]
		}
	}

	// Navigate the map
	foundZZZ := false
	leftRightInstructionsCounter := 0
	stepsCounter := 0
	currentInstruction := "AAA"

	for {
		if foundZZZ {
			break
		}

		currentLeftOrRightInstruction := getLeftOrRightInstruction(leftRightInstructions, &leftRightInstructionsCounter)

		for _, r := range directions {
			if currentInstruction == r[:3] {
				leftInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((strings.Split(r, "=")[1]), "(", ""), ")", ""), ",", ""))[0]
				rightInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((strings.Split(r, "=")[1]), "(", ""), ")", ""), ",", ""))[1]

				if currentLeftOrRightInstruction == "L" {
					currentInstruction = leftInstruction
				}
				if currentLeftOrRightInstruction == "R" {
					currentInstruction = rightInstruction
				}
				break
			}
		}
		stepsCounter++
		if currentInstruction == "ZZZ" {
			foundZZZ = true
		}
	}

	fmt.Println("Total steps: " + strconv.Itoa(stepsCounter))

	fmt.Println("*** End part A ***")
}

func processDay8B(inputLines []string) {
	fmt.Println("*** Start part B ***")
	leftRightInstructions := ""
	directions := make([]string, 0)
	directionsMap := make(map[string]string)

	// Read in the input file and split out the important pieces
	for i := 0; i < len(inputLines); i++ {

		inputLine := inputLines[i]
		if len(inputLine) == 0 {
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "=") {
			directions = append(directions, inputLine)
			myKey := inputLine[:3]
			myValue := strings.Split(inputLine, "=")[1]
			directionsMap[myKey] = myValue
			continue
		} else {
			leftRightInstructions = inputLines[i]
		}
	}

	//sort.Strings(directions)

	for _, r := range directions {
		fmt.Println(r)
	}

	// Navigate the map
	foundZZZ := false
	leftRightInstructionsCounter := 0
	stepsCounter := 0
	currentInstructions := getInitialInstructions(directions)
	currentInstructions = currentInstructions[5:6]

	foundXTimes := 0

	for {
		if foundZZZ && foundXTimes > 5 {
			break
		}

		currentLeftOrRightInstruction := getLeftOrRightInstruction(leftRightInstructions, &leftRightInstructionsCounter)
		if stepsCounter == 0 {
			printCurrentInstructions(currentInstructions, stepsCounter, currentLeftOrRightInstruction, leftRightInstructionsCounter)
		}
		for c := 0; c < len(currentInstructions); c++ {

			//for _, r := range directions {
			///	if strings.Contains(r, currentInstructions[c]) && currentInstructions[c] == r[:3] {
			//		leftInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((strings.Split(r, "=")[1]), "(", ""), ")", ""), ",", ""))[0]
			//		rightInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((strings.Split(r, "=")[1]), "(", ""), ")", ""), ",", ""))[1]
			leftInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((directionsMap[currentInstructions[c]]), "(", ""), ")", ""), ",", ""))[0]
			rightInstruction := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll((directionsMap[currentInstructions[c]]), "(", ""), ")", ""), ",", ""))[1]
			if currentLeftOrRightInstruction == "L" {
				currentInstructions[c] = leftInstruction
			}
			if currentLeftOrRightInstruction == "R" {
				currentInstructions[c] = rightInstruction
			}
			//break
			//	}
			//}
		}

		/*if stepsCounter%1000000 == 0 {
			printCurrentInstructions(currentInstructions, stepsCounter)
		}*/
		//printCurrentInstructions(currentInstructions, stepsCounter, currentLeftOrRightInstruction, leftRightInstructionsCounter)
		if areCurrentInstructionsAllZs(currentInstructions) {
			foundZZZ = true
			foundXTimes++
			printCurrentInstructions(currentInstructions, stepsCounter, currentLeftOrRightInstruction, leftRightInstructionsCounter)
		}
		////if anyZs(currentInstructions) || stepsCounter == 1 {
		////	printCurrentInstructions(currentInstructions, stepsCounter, currentLeftOrRightInstruction, leftRightInstructionsCounter)
		//}
		stepsCounter++
	}

	fmt.Println("Total steps: " + strconv.Itoa(stepsCounter))

	fmt.Printf("*** End part B ***")
}

func getInitialInstructions(directions []string) []string {
	returnVal := make([]string, 0)

	for _, r := range directions {
		if r[2:3] == "A" {
			returnVal = append(returnVal, r[:3])
		}
	}

	return returnVal
}

func areCurrentInstructionsAllZs(currentInstructions []string) bool {
	returnVal := true

	for _, c := range currentInstructions {
		if c[2:3] != "Z" {
			return false
		}
	}

	return returnVal
}

func anyZs(currentInstructions []string) bool {
	returnVal := false

	for _, c := range currentInstructions {
		if c[2:3] == "Z" {
			return true
		}
	}

	return returnVal
}

func printCurrentInstructions(currentInstructions []string, stepCounter int, currRightLeft string, currRightLeftCounter int) {
	fmt.Println("Current Instructions(" + strconv.Itoa(stepCounter) + "): " + strings.Join(currentInstructions, ", ") + " :: " + currRightLeft + ", " + strconv.Itoa(currRightLeftCounter))
}
