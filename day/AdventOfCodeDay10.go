package day

import (
	"fmt"
	"strconv"
)

type mapPosition struct {
	x     int
	y     int
	prevx int
	prevy int
}

func ProcessDay10(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	if aocDay.Part == "A" {
		processDay10A(inputLines)
	} else {
		processDay10B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay10A(inputLines []string) {
	fmt.Println("*** Start part A ***")

	pipeMap := make([][]rune, 0)
	for i := range pipeMap {
		pipeMap[i] = make([]rune, 0)
	}

	fmt.Println("Input: ")
	var startPosition mapPosition = mapPosition{
		x: -1,
		y: -1}

	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		tempXAxis := make([]rune, 0)
		for j, r := range inputLine {
			tempXAxis = append(tempXAxis, r)

			if r == rune('S') {
				fmt.Println("Found the start position!  Row: " + strconv.Itoa(i) + ", Column: " + strconv.Itoa(j))
				startPosition = mapPosition{
					x:     j,
					y:     i,
					prevx: -1,
					prevy: -1}
			}
		}
		pipeMap = append(pipeMap, tempXAxis)

		//fmt.Println(inputLine)
	}
	numberOfMovesUntilTheEnd := solveMapAndGetHalfwayPoint(pipeMap, startPosition)

	printPipeMap(pipeMap)
	fmt.Println("Total number of moves: " + strconv.Itoa(numberOfMovesUntilTheEnd))
	fmt.Println("*** End part A ***")
}

func printPipeMap(pipeMap [][]rune) {
	for y := 0; y < len(pipeMap); y++ {
		for x := 0; x < len(pipeMap[y]); x++ {
			fmt.Print("" + string(pipeMap[y][x]))
		}
		fmt.Println("")
	}
}

func solveMapAndGetHalfwayPoint(pipeMap [][]rune, currentPosition mapPosition) int {
	nextPosition := decideOnNextPosition(pipeMap, currentPosition)
	currentRune := pipeMap[nextPosition.y][nextPosition.x]

	if currentRune == rune('S') && currentPosition.prevx != -1 && currentPosition.prevy != -1 {
		// We found the end!
		return 0
	}

	return solveMapAndGetHalfwayPoint(pipeMap, nextPosition) + 1

}

func decideOnNextPosition(pipeMap [][]rune, currentPosition mapPosition) mapPosition {
	printMapPosition("decideOnNextPosition: ", currentPosition)
	isCanGoUp := getCanGoUp(pipeMap, currentPosition)
	isCanGoRight := getCanGoRight(pipeMap, currentPosition)
	isCanGoDown := getCanGoDown(pipeMap, currentPosition)
	isCanGoLeft := getCanGoLeft(pipeMap, currentPosition)

	if isCanGoUp {

		tempPosition := mapPosition{
			x:     currentPosition.x,
			y:     currentPosition.y - 1,
			prevx: currentPosition.x,
			prevy: currentPosition.y}
		if !(tempPosition.x == currentPosition.prevx && tempPosition.y == currentPosition.prevy) {
			fmt.Println("Going Up!")
			return tempPosition
		}
	}
	if isCanGoRight {

		tempPosition := mapPosition{
			x:     currentPosition.x + 1,
			y:     currentPosition.y,
			prevx: currentPosition.x,
			prevy: currentPosition.y}
		if !(tempPosition.x == currentPosition.prevx && tempPosition.y == currentPosition.prevy) {
			fmt.Println("Going Right!")
			return tempPosition
		}
	}
	if isCanGoDown {

		tempPosition := mapPosition{
			x:     currentPosition.x,
			y:     currentPosition.y + 1,
			prevx: currentPosition.x,
			prevy: currentPosition.y}
		if !(tempPosition.x == currentPosition.prevx && tempPosition.y == currentPosition.prevy) {
			fmt.Println("Going Down!")
			return tempPosition
		}
	}
	if isCanGoLeft {

		tempPosition := mapPosition{
			x:     currentPosition.x - 1,
			y:     currentPosition.y,
			prevx: currentPosition.x,
			prevy: currentPosition.y}
		if !(tempPosition.x == currentPosition.prevx && tempPosition.y == currentPosition.prevy) {
			fmt.Println("Going Left!")
			return tempPosition
		}
	}

	fmt.Println("ERROR!")
	panic("Error found!")

	return mapPosition{
		0, 0, 0, 0}
}

func getCanGoUp(pipeMap [][]rune, currentPosition mapPosition) bool {
	validUpCharacters := []rune{'|', 'F', '7', 'S'}
	validCurrentCharacters := []rune{'S', 'L', 'J', '|'}
	currentRune := pipeMap[currentPosition.y][currentPosition.x]
	isValidCurrentCharacter := runeDoesContain(validCurrentCharacters, currentRune)
	if currentPosition.y == 0 {
		return false
	}
	isValidCharacterAbove := runeDoesContain(validUpCharacters, pipeMap[currentPosition.y-1][currentPosition.x])

	return isValidCurrentCharacter && isValidCharacterAbove
}

func getCanGoRight(pipeMap [][]rune, currentPosition mapPosition) bool {
	validRightCharacters := []rune{'-', '7', 'J', 'S'}
	validCurrentCharacters := []rune{'S', 'F', 'L', '-'}
	currentRuneRow := pipeMap[currentPosition.y]
	currentRune := pipeMap[currentPosition.y][currentPosition.x]
	isValidCurrentCharacter := runeDoesContain(validCurrentCharacters, currentRune)
	if currentPosition.x == (len(currentRuneRow) - 1) {
		return false
	}
	isValidCharacterToTheRight := runeDoesContain(validRightCharacters, pipeMap[currentPosition.y][currentPosition.x+1])

	fmt.Println("isValidCurrentCharacter: " + strconv.FormatBool(isValidCurrentCharacter))
	fmt.Println("isValidCharacterToTheRight: " + strconv.FormatBool(isValidCharacterToTheRight))

	return isValidCurrentCharacter && isValidCharacterToTheRight
}

func getCanGoDown(pipeMap [][]rune, currentPosition mapPosition) bool {
	validDownCharacters := []rune{'|', 'L', 'J', 'S'}
	validCurrentCharacters := []rune{'S', 'F', '7', '|'}
	currentRune := pipeMap[currentPosition.y][currentPosition.x]
	isValidCurrentCharacter := runeDoesContain(validCurrentCharacters, currentRune)
	if currentPosition.y == (len(pipeMap) - 1) {
		return false
	}
	isValidCharacterDown := runeDoesContain(validDownCharacters, pipeMap[currentPosition.y+1][currentPosition.x])

	return isValidCurrentCharacter && isValidCharacterDown
}

func getCanGoLeft(pipeMap [][]rune, currentPosition mapPosition) bool {
	validLeftCharacters := []rune{'-', 'F', 'L', 'S'}
	validCurrentCharacters := []rune{'S', '7', 'J', '-'}
	currentRune := pipeMap[currentPosition.y][currentPosition.x]
	isValidCurrentCharacter := runeDoesContain(validCurrentCharacters, currentRune)
	if currentPosition.x == 0 {
		return false
	}
	isValidCharacterLeft := runeDoesContain(validLeftCharacters, pipeMap[currentPosition.y][currentPosition.x-1])

	return isValidCurrentCharacter && isValidCharacterLeft
}

func runeDoesContain(validCharacterList []rune, checkCharacter rune) bool {
	for i := 0; i < len(validCharacterList); i++ {
		if validCharacterList[i] == checkCharacter {
			return true
		}
	}
	return false
}

func printMapPosition(prefix string, currentPosition mapPosition) {
	fmt.Println(prefix + " Current position: y=" + strconv.Itoa(currentPosition.y) + ", x=" + strconv.Itoa(currentPosition.x))
}

func processDay10B(inputLines []string) {
	fmt.Println("*** Start part B ***")

	pipeMap := make([][]rune, 0)
	for i := range pipeMap {
		pipeMap[i] = make([]rune, 0)
	}

	fmt.Println("Input: ")
	var startPosition mapPosition = mapPosition{
		x: -1,
		y: -1}

	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		tempXAxis := make([]rune, 0)
		for j, r := range inputLine {
			tempXAxis = append(tempXAxis, r)

			if r == rune('S') {
				fmt.Println("Found the start position!  Row: " + strconv.Itoa(i) + ", Column: " + strconv.Itoa(j))
				startPosition = mapPosition{
					x:     j,
					y:     i,
					prevx: -1,
					prevy: -1}
			}
		}
		pipeMap = append(pipeMap, tempXAxis)

		//fmt.Println(inputLine)
	}
	numberOfMovesUntilTheEnd := solveMapAndGetHalfwayPointer(&pipeMap, startPosition)

	printPipeMap(pipeMap)
	fmt.Println("Total number of moves: " + strconv.Itoa(numberOfMovesUntilTheEnd))
	fmt.Printf("*** End part B ***")
}

func solveMapAndGetHalfwayPointer(pipeMap *[][]rune, currentPosition mapPosition) int {
	nextPosition := decideOnNextPosition((*pipeMap), currentPosition)
	currentRune := (*pipeMap)[nextPosition.y][nextPosition.x]

	if currentRune == rune('S') && currentPosition.prevx != -1 && currentPosition.prevy != -1 {
		// We found the end!
		(*pipeMap)[currentPosition.y][currentPosition.x] = rune('*')
		return 0
	}
	if (*pipeMap)[currentPosition.y][currentPosition.x] != rune('S') {
		(*pipeMap)[currentPosition.y][currentPosition.x] = rune('*')
	}

	return solveMapAndGetHalfwayPointer(pipeMap, nextPosition) + 1

}
