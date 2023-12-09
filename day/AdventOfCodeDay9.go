package day

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProcessDay9(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	if aocDay.Part == "A" {
		processDay9A(inputLines)
	} else {
		processDay9B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay9A(inputLines []string) {
	fmt.Println("Input: ")
	total := 0
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		listOfNumbersStr := strings.Fields(inputLine)
		listOfNumbers := ConvertStrSliceToIntSlice(listOfNumbersStr)
		printIntSlice(listOfNumbers)

		nextNum := calculateNextValue(listOfNumbers)
		//fmt.Println("Found next num: " + strconv.Itoa(nextNum))
		nextVal := listOfNumbers[len(listOfNumbers)-1] + nextNum
		printIntSliceAndValue(listOfNumbers, nextVal)
		fmt.Println("Found next value: " + strconv.Itoa(nextVal))
		total += nextVal
		fmt.Println("Total: " + strconv.Itoa(total))

		// Pause for keyboard input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Press enter to continue")
		reader.ReadString('\n')

	}

	fmt.Println("Total: " + strconv.Itoa(total))
}

func calculateNextValue(intSlice []int) int {
	differenceInNumbers := make([]int, 0)
	//printIntSlice(intSlice)

	for i := 0; i < len(intSlice)-1; i++ {
		//tempVal := distanceBetweenTwoNumbers(intSlice[i+1], intSlice[i])
		tempVal := intSlice[i+1] - intSlice[i]
		differenceInNumbers = append(differenceInNumbers, tempVal)
	}

	if allNumbersAreZero(differenceInNumbers) {
		printIntSliceAndValue(differenceInNumbers, 0)
		return 0
	}
	tempVal := calculateNextValue(differenceInNumbers)
	//printIntSliceAndValue(differenceInNumbers, tempVal)

	// Pause for keyboard input
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Press enter to continue - 2 ")
	//reader.ReadString('\n')

	//fmt.Println("returning value: " + strconv.Itoa(tempVal+differenceInNumbers[len(differenceInNumbers)-1]))
	returnVal := tempVal + differenceInNumbers[len(differenceInNumbers)-1]
	printIntSliceAndValue(differenceInNumbers, returnVal)
	return returnVal
}

func totalIntSlice(intSlice []int) int {
	returnVal := 0

	for _, r := range intSlice {
		returnVal += r
	}

	return returnVal
}

func allNumbersAreZero(intSlice []int) bool {

	for _, r := range intSlice {
		if r != 0 {
			return false
		}
	}
	return true
}

func printIntSlice(intSlice []int) {
	strSlice := make([]string, 0)
	for _, r := range intSlice {
		tempStr := strconv.Itoa(r)
		strSlice = append(strSlice, tempStr)
	}
	fmt.Println(strings.Join(strSlice, ", "))
}

func printIntSliceAndValue(intSlice []int, i int) {
	strSlice := make([]string, 0)
	for _, r := range intSlice {
		tempStr := strconv.Itoa(r)
		strSlice = append(strSlice, tempStr)
	}
	//lastVal := intSlice[len(intSlice)-1] + i
	fmt.Println(strings.Join(strSlice, ", ") + ", (" + strconv.Itoa(i) + ")")
}

func processDay9B(inputLines []string) {
	fmt.Println("Input: ")
	total := 0
	for i := 0; i < len(inputLines); i++ {

		inputLine := inputLines[i]
		listOfNumbersStr := strings.Fields(inputLine)
		for i, j := 0, len(listOfNumbersStr)-1; i < j; i, j = i+1, j-1 {
			listOfNumbersStr[i], listOfNumbersStr[j] = listOfNumbersStr[j], listOfNumbersStr[i]
		}
		listOfNumbers := ConvertStrSliceToIntSlice(listOfNumbersStr)
		printIntSlice(listOfNumbers)

		nextNum := calculateNextValue(listOfNumbers)
		//fmt.Println("Found next num: " + strconv.Itoa(nextNum))
		nextVal := listOfNumbers[len(listOfNumbers)-1] + nextNum
		printIntSliceAndValue(listOfNumbers, nextVal)
		fmt.Println("Found next value: " + strconv.Itoa(nextVal))
		total += nextVal
		fmt.Println("Total: " + strconv.Itoa(total))

		// Pause for keyboard input
		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Press enter to continue")
		//reader.ReadString('\n')

	}

	fmt.Println("Total: " + strconv.Itoa(total))
}
