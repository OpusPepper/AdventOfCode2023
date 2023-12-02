package main

import (
	"fmt"
	"strconv"
	"strings"
)

type groupOfBlocks struct {
	numberOfBlocks int
	colorOfBlocks  string
}

type bag struct {
	red   groupOfBlocks
	blue  groupOfBlocks
	green groupOfBlocks
}

func getGame(inputLine string) int {
	v := strings.Split(inputLine, ":")
	fields := strings.Fields(v[0])
	fmt.Println(fields[1])
	val, err := strconv.Atoi(fields[1])
	check(err)
	return val
}

func testBag(numberOfBlocks int, colorOfBlocks string, testBag bag) bool {
	if colorOfBlocks == "blue" {
		return numberOfBlocks <= testBag.blue.numberOfBlocks
	}
	if colorOfBlocks == "red" {
		return numberOfBlocks <= testBag.red.numberOfBlocks
	}
	if colorOfBlocks == "green" {
		return numberOfBlocks <= testBag.green.numberOfBlocks
	}
	return false
}

func getPieces(inputLine string, solutionBag bag) bool {
	returnBool := true
	v := strings.Split(inputLine, ":")
	//fmt.Println("v1: " + v[1])
	s := strings.Split(v[1], ";")
	for i := 0; i < len(s); i++ {
		//fmt.Println("s[" + strconv.Itoa(i) + "]: " + s[i])
		c := strings.Split(s[i], ",")
		for j := 0; j < len(c); j++ {
			//fmt.Println("c[" + strconv.Itoa(j) + "]: " + c[j])
			f := strings.Fields(c[j])
			numberOfColorBlocks, _ := strconv.Atoi(f[0])
			colorOfBlocks := f[1]
			isBlocksOk := testBag(numberOfColorBlocks, colorOfBlocks, solutionBag)
			if !isBlocksOk {
				fmt.Println("This game doesn't work, failing")
				return false
			}
			returnBool = returnBool && isBlocksOk
		}
	}
	return returnBool
}

func getMinimumPieces(inputLine string) bag {
	v := strings.Split(inputLine, ":")
	//fmt.Println("v1: " + v[1])
	blueBlocks := groupOfBlocks{0, "blue"}
	redBlocks := groupOfBlocks{0, "red"}
	greenBlocks := groupOfBlocks{0, "green"}

	s := strings.Split(v[1], ";")
	for i := 0; i < len(s); i++ {
		//fmt.Println("s[" + strconv.Itoa(i) + "]: " + s[i])
		c := strings.Split(s[i], ",")
		for j := 0; j < len(c); j++ {
			//fmt.Println("c[" + strconv.Itoa(j) + "]: " + c[j])
			f := strings.Fields(c[j])
			numberOfColorBlocks, _ := strconv.Atoi(f[0])
			colorOfBlocks := f[1]
			if colorOfBlocks == redBlocks.colorOfBlocks {
				if numberOfColorBlocks > redBlocks.numberOfBlocks {
					redBlocks.numberOfBlocks = numberOfColorBlocks
				}
			}
			if colorOfBlocks == greenBlocks.colorOfBlocks {
				if numberOfColorBlocks > greenBlocks.numberOfBlocks {
					greenBlocks.numberOfBlocks = numberOfColorBlocks
				}
			}
			if colorOfBlocks == blueBlocks.colorOfBlocks {
				if numberOfColorBlocks > blueBlocks.numberOfBlocks {
					blueBlocks.numberOfBlocks = numberOfColorBlocks
				}
			}
		}
	}
	return bag{redBlocks, blueBlocks, greenBlocks}
}

func processDay2A(inputLines []string) {
	fmt.Println("*** Day 2 part A ***")
	total := 0

	//Bag has this many blocks:
	blueBlocks := groupOfBlocks{14, "blue"}
	redBlocks := groupOfBlocks{12, "red"}
	greenBlocks := groupOfBlocks{13, "green"}
	solutionBag := bag{
		redBlocks,
		blueBlocks,
		greenBlocks}

	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		gameDayInt := getGame(inputLine)
		fmt.Println("Game day: " + strconv.Itoa(gameDayInt))
		isGamePassed := getPieces(inputLine, solutionBag)
		if isGamePassed {
			total = total + gameDayInt
		}
	}
	fmt.Println("Total for games: " + strconv.Itoa(total))
	fmt.Println("*** End processing ***")
}

func processDay2B(inputLines []string) {
	fmt.Println("*** Day 2 part B ***")
	total := 0

	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		gameDayInt := getGame(inputLine)
		fmt.Println("Game day: " + strconv.Itoa(gameDayInt))
		potentialSolutionBag := getMinimumPieces(inputLine)
		powerOfCubes := potentialSolutionBag.blue.numberOfBlocks *
			potentialSolutionBag.green.numberOfBlocks *
			potentialSolutionBag.red.numberOfBlocks
		total = total + powerOfCubes
	}
	fmt.Println("Total for games: " + strconv.Itoa(total))
	fmt.Println("*** End processing ***")
}
