package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findFirstNumber(line string) string {
	for i := 0; i < len(line); i++ {
		evalChar := string(line[i])
		_, err := strconv.Atoi(evalChar)
		if err == nil {
			//fmt.Println("Found first number! -> " + evalChar)
			return evalChar
		}
	}
	return "0"
}

func findLastNumber(line string) string {
	for i := (len(line) - 1); i >= 0; i-- {
		evalChar := string(line[i])
		_, err := strconv.Atoi(evalChar)
		if err == nil {
			//fmt.Println("Found last number! -> " + evalChar)
			return evalChar
		}
	}
	return "0"
}

func convertInput(line string) string {
	lowerLine := strings.ToLower(line)
	if strings.Contains(lowerLine, "twone") {
		lowerLine = strings.ReplaceAll(lowerLine, "twone", "21")
	}
	if strings.Contains(lowerLine, "oneight") {
		lowerLine = strings.ReplaceAll(lowerLine, "oneight", "18")
	}
	if strings.Contains(lowerLine, "threeight") {
		lowerLine = strings.ReplaceAll(lowerLine, "threeight", "38")
	}
	if strings.Contains(lowerLine, "fiveight") {
		lowerLine = strings.ReplaceAll(lowerLine, "fiveight", "58")
	}
	if strings.Contains(lowerLine, "nineight") {
		lowerLine = strings.ReplaceAll(lowerLine, "nineight", "98")
	}
	if strings.Contains(lowerLine, "eightwo") {
		lowerLine = strings.ReplaceAll(lowerLine, "eightwo", "82")
	}
	if strings.Contains(lowerLine, "eighthree") {
		lowerLine = strings.ReplaceAll(lowerLine, "eighthree", "83")
	}
	if strings.Contains(lowerLine, "one") {
		lowerLine = strings.ReplaceAll(lowerLine, "one", "1")
	}
	if strings.Contains(lowerLine, "two") {
		lowerLine = strings.ReplaceAll(lowerLine, "two", "2")
	}
	if strings.Contains(lowerLine, "three") {
		lowerLine = strings.ReplaceAll(lowerLine, "three", "3")
	}
	if strings.Contains(lowerLine, "four") {
		lowerLine = strings.ReplaceAll(lowerLine, "four", "4")
	}
	if strings.Contains(lowerLine, "five") {
		lowerLine = strings.ReplaceAll(lowerLine, "five", "5")
	}
	if strings.Contains(lowerLine, "six") {
		lowerLine = strings.ReplaceAll(lowerLine, "six", "6")
	}
	if strings.Contains(lowerLine, "seven") {
		lowerLine = strings.ReplaceAll(lowerLine, "seven", "7")
	}
	if strings.Contains(lowerLine, "eight") {
		lowerLine = strings.ReplaceAll(lowerLine, "eight", "8")
	}
	if strings.Contains(lowerLine, "nine") {
		lowerLine = strings.ReplaceAll(lowerLine, "nine", "9")
	}
	return lowerLine
}

func processDay1A(inputLines []string) {
	fmt.Println("*** Day 1 part A ***")
	total := 0

	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		firstNumberString := findFirstNumber(inputLines[i])
		lastNumberString := findLastNumber(inputLines[i])

		finalNumberString := firstNumberString + lastNumberString

		finalNumber, err := strconv.Atoi(finalNumberString)
		check(err)

		total = total + finalNumber

		fmt.Println("Number " + strconv.Itoa(i) + ": " + finalNumberString + ", total: " + strconv.Itoa(total) + ", line: " + inputLines[i])
	}
	fmt.Println()
}

func processDay1B(inputLines []string) {
	fmt.Println("*** Day 1 part B ***")
	total := 0

	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		//fmt.Println("Line " + strconv.Itoa(i) + ": " + inputLines[i])

		convertedLine := convertInput(inputLines[i])

		firstNumberString := findFirstNumber(convertedLine)
		lastNumberString := findLastNumber(convertedLine)

		finalNumberString := firstNumberString + lastNumberString

		finalNumber, err := strconv.Atoi(finalNumberString)
		check(err)

		total = total + finalNumber

		fmt.Println("Number " + strconv.Itoa(i) + ": " + finalNumberString + ", total: " + strconv.Itoa(total) + ", line: " + convertedLine)
	}
	fmt.Println()
}
