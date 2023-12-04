package main

import (
	"fmt"
	"strconv"
	"strings"
)

type card struct {
	cardNumber          int
	winningNumbers      []string
	selectedNumbers     []string
	numberOfMatches     int
	numberOfCardsToPlay int
}

func getScratchNumbersA(inputLine string, inRow int) int {
	splitOnColon := strings.Split(inputLine, ":")
	cardNumber := strings.Fields(splitOnColon[0])[1]
	allNumbers := strings.Split(splitOnColon[1], "|")
	winningNumbers := strings.Fields(allNumbers[0])
	selectedNumbers := strings.Fields(allNumbers[1])
	numberOfMatches := 0

	fmt.Println("SplitOnColon: " + splitOnColon[0])
	//numberString := ""
	for w := 0; w < len(winningNumbers); w++ {
		for s := 0; s < len(selectedNumbers); s++ {
			if winningNumbers[w] == selectedNumbers[s] {
				fmt.Println("We found a winning number Card " + cardNumber + " Number " + winningNumbers[w])
				numberOfMatches = numberOfMatches + 1
			}
		}
	}
	cardNumberInt, _ := strconv.Atoi(cardNumber)

	currentCard := card{
		cardNumberInt,
		winningNumbers,
		selectedNumbers,
		numberOfMatches,
		1}
	fmt.Println("Found " + strconv.Itoa(currentCard.numberOfMatches) + " matches!!")
	myReturnVal := doubleMyXNumberYTimes(1, numberOfMatches)
	if currentCard.numberOfMatches == 0 {
		myReturnVal = 0
	}
	fmt.Println("doubleMyXNumberYTimes: " + strconv.Itoa(myReturnVal))

	return myReturnVal
}

func doubleMyXNumberYTimes(x int, y int) int {
	returnVal := x
	for i := 0; i < (y - 1); i++ {
		returnVal = returnVal * 2
	}
	return returnVal
}

func getScratchNumbersB(inputLine string, inRow int) card {
	splitOnColon := strings.Split(inputLine, ":")
	cardNumber := strings.Fields(splitOnColon[0])[1]
	allNumbers := strings.Split(splitOnColon[1], "|")
	winningNumbers := strings.Fields(allNumbers[0])
	selectedNumbers := strings.Fields(allNumbers[1])
	numberOfMatches := 0

	fmt.Println("SplitOnColon: " + splitOnColon[0])
	//numberString := ""
	for w := 0; w < len(winningNumbers); w++ {
		for s := 0; s < len(selectedNumbers); s++ {
			if winningNumbers[w] == selectedNumbers[s] {
				//fmt.Println("We found a winning number Card " + cardNumber + " Number " + winningNumbers[w])
				numberOfMatches = numberOfMatches + 1
			}
		}
	}
	cardNumberInt, _ := strconv.Atoi(cardNumber)

	currentCard := card{
		cardNumberInt,
		winningNumbers,
		selectedNumbers,
		numberOfMatches,
		1}
	//fmt.Println("Found " + strconv.Itoa(currentCard.numberOfMatches) + " matches!!")
	//myReturnVal := doubleMyXNumberYTimes(1, numberOfMatches)
	//if currentCard.numberOfMatches == 0 {
	//		myReturnVal = 0
	//	}
	//fmt.Println("doubleMyXNumberYTimes: " + strconv.Itoa(myReturnVal))

	return currentCard
}

func processDay4(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.day) + " part " + aocDay.part + " ***")

	//initialize objects needed

	if aocDay.part == "A" {
		processDay4A(inputLines)
	} else {
		processDay4B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay4A(inputLines []string) {
	fmt.Println("Input: ")
	total := 0
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		total = total + getScratchNumbersA(inputLine, i)
		//i = len(inputLines)
	}

	fmt.Println("Total: " + strconv.Itoa(total))
}

func processDay4B(inputLines []string) {
	fmt.Println("Input: ")
	total := 0
	var newCard card
	var allCards []card

	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		newCard = getScratchNumbersB(inputLine, i)
		//i = len(inputLines)
		allCards = append(allCards, newCard)
	}

	for i := 0; i < len(allCards); i++ {
		isWinningCard := allCards[i].numberOfMatches > 0
		numberOfPlays := allCards[i].numberOfCardsToPlay
		if isWinningCard {
			for j := 0; j < numberOfPlays; j++ {
				updateAllCards(&allCards, allCards[i].cardNumber, allCards[i].numberOfMatches)
			}
		}
	}

	total = 0
	for i := 0; i < len(allCards); i++ {
		//isWinningCard := allCards[i].numberOfMatches > 0
		//fmt.Println("Card[" + strconv.Itoa(allCards[i].cardNumber) + "] Number of cards played: " + strconv.Itoa(allCards[i].numberOfCardsToPlay))
		//fmt.Println("isWinning: " + strconv.FormatBool(isWinningCard))

		total = total + allCards[i].numberOfCardsToPlay
	}

	fmt.Println("Total: " + strconv.Itoa(total))
}

func updateAllCards(allCards *[]card, cardNumber int, numberOfMatches int) {
	for i := 0; i < numberOfMatches; i++ {

		if (cardNumber + i) < len(*allCards) {
			(*allCards)[cardNumber+i].numberOfCardsToPlay = (*allCards)[cardNumber+i].numberOfCardsToPlay + 1
		}
	}
}
