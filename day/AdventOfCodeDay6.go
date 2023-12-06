package day

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStrSliceToIntSlice(inputSlice []string) []int {
	var outputSlice []int

	for i := 0; i < len(inputSlice); i++ {
		sliceVal, err := strconv.Atoi(inputSlice[i])
		check(err)
		outputSlice = append(outputSlice, sliceVal)
	}

	return outputSlice
}

func ProcessDay6(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	//initialize objects needed

	if aocDay.Part == "A" {
		processDay6A(inputLines)
	} else {
		processDay6B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay6A(inputLines []string) {
	fmt.Println("Input: ")

	//for i := 0; i < len(inputLines); i++ {
	//inputLine := inputLines[i]

	timeSliceStr := strings.Fields(inputLines[0])[1:]
	distanceSliceStr := strings.Fields(inputLines[1])[1:]

	timeSliceInt := ConvertStrSliceToIntSlice(timeSliceStr)
	distanceSliceInt := ConvertStrSliceToIntSlice(distanceSliceStr)
	//}
	grandTotal := 0

	for i := 0; i < len(timeSliceInt); i++ {
		currentTime := timeSliceInt[i]
		currentDistanceRecord := distanceSliceInt[i]
		totalWins := 0
		for t := 1; t <= currentTime; t++ {
			pressedFor := t
			wentFor := pressedFor * (currentTime - pressedFor)
			if wentFor > currentDistanceRecord {
				totalWins = totalWins + 1
			}
		}
		fmt.Println("Found " + strconv.Itoa(totalWins) + " wins")
		if grandTotal == 0 && totalWins > 0 {
			grandTotal = totalWins
			continue
		}
		if totalWins != 0 {
			grandTotal = grandTotal * totalWins
		}

	}

	fmt.Println("grandTotal: " + strconv.Itoa(grandTotal))
}

func processDay6B(inputLines []string) {
	fmt.Println("Input: ")

	//for i := 0; i < len(inputLines); i++ {
	//inputLine := inputLines[i]

	timeSliceStr := strings.Join(strings.Fields(inputLines[0])[1:], "")
	distanceSliceStr := strings.Join(strings.Fields(inputLines[1])[1:], "")

	timeSliceInt, err := strconv.Atoi(timeSliceStr)
	check(err)
	distanceSliceInt, err := strconv.Atoi(distanceSliceStr)
	//}

	//for i := 0; i < len(timeSliceInt); i++ {
	currentTime := timeSliceInt
	currentDistanceRecord := distanceSliceInt
	totalWins := 0
	for t := 1; t <= currentTime; t++ {
		pressedFor := t
		wentFor := pressedFor * (currentTime - pressedFor)
		if wentFor > currentDistanceRecord {
			totalWins = totalWins + 1
		}
	}
	fmt.Println("Found " + strconv.Itoa(totalWins) + " wins")

}
