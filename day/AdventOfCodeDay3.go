package day

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AdventOfCodeDay struct {
	Day          int
	Part         string
	FileOverride string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type machinePart struct {
	partNumber int
	partString string
	row        int
	column     int
}

type machineSymbol struct {
	partSymbol string
	row        int
	column     int
	partCount  int
	gearRatio  int
}

type fullLineParts struct {
	lineMachineParts   *[]machinePart
	lineMachineSymbols *[]machineSymbol
}

func getNumberStringsOnLine(inputLine string) []string {
	numbersOrSymbols := strings.Fields(strings.ReplaceAll(inputLine, ".", " "))
	var returnNumberStrings []string
	for i := 0; i < len(numbersOrSymbols); i++ {
		_, err := strconv.Atoi(numbersOrSymbols[i])
		if err == nil {
			returnNumberStrings = append(returnNumberStrings, numbersOrSymbols[i])
			continue
		}
	}

	return returnNumberStrings
}

func getNumbersAndSymbols(inputLine string, inRow int, machineParts *[]machinePart, machineSymbols *[]machineSymbol) {
	//numberString := ""
	var currentPart machinePart
	for i := 0; i < len(inputLine); i++ {
		runeInputLine := string([]rune(inputLine))
		runeToString := string(runeInputLine[i])
		asciiCharVal := int(runeInputLine[i])
		isPreviousCharANumber := false

		if i > 0 {
			prevAsciiCharVal := int(runeInputLine[i-1])
			isPreviousCharANumber = 48 <= prevAsciiCharVal && prevAsciiCharVal <= 57
		}
		if isPeriod := runeInputLine[i] == '.'; isPeriod {
			if currentPart != (machinePart{}) {
				runeToInt, _ := strconv.Atoi(currentPart.partString)
				currentPart.partNumber = runeToInt
				*machineParts = append(*machineParts, currentPart)
				currentPart = machinePart{}
			}
			continue
		}
		//fmt.Println("Char is: " + string(runeInputLine[i]))

		//fmt.Println("ASCII is: " + string(runeInputLine[i]))
		isNumber := 48 <= asciiCharVal && asciiCharVal <= 57
		if isNumber {
			if isPreviousCharANumber {
				currentPart.partString = currentPart.partString + runeToString
				continue
			}

			// Save number to a new structure
			currentPart = machinePart{
				partString: runeToString,
				row:        inRow,
				column:     i}
			continue
		}

		// Otherwise, it's a symbol
		if currentPart != (machinePart{}) {
			runeToInt, _ := strconv.Atoi(currentPart.partString)
			currentPart.partNumber = runeToInt
			*machineParts = append(*machineParts, currentPart)
			currentPart = machinePart{}
		}
		s := machineSymbol{
			partSymbol: runeToString,
			row:        inRow,
			column:     i,
			partCount:  0,
			gearRatio:  0}
		*machineSymbols = append(*machineSymbols, s)
	}
	//fmt.Println("> Number of machine parts: " + strconv.Itoa(len(*machineParts)))
	//fmt.Println("> Number of symbols: " + strconv.Itoa(len(*machineSymbols)))

	if currentPart != (machinePart{}) {
		runeToInt, _ := strconv.Atoi(currentPart.partString)
		currentPart.partNumber = runeToInt
		*machineParts = append(*machineParts, currentPart)
		currentPart = machinePart{}
	}

}

func checkPartsTouchingSymbols(machineParts []machinePart, machineSymbols []machineSymbol) {
	total := 0

	for p := 0; p < len(machineParts); p++ {
		currentPart := machineParts[p]
		//fmt.Println("Current part is: " + currentPart.partString)

		for s := 0; s < len(machineSymbols); s++ {
			currentSymbol := &machineSymbols[s]

			rowCheck := math.Abs(float64(currentPart.row) - float64(currentSymbol.row))

			isAboveOrBelow := rowCheck == 1 || rowCheck == 0
			isNextTo := (currentPart.column-1 <= currentSymbol.column) &&
				((currentPart.column + len(currentPart.partString)) >= currentSymbol.column)

			if isAboveOrBelow && isNextTo {
				total = total + currentPart.partNumber
				currentSymbol.partCount = currentSymbol.partCount + 1
				//fmt.Println("Count: " + strconv.Itoa(currentSymbol.partCount))
				if currentSymbol.partCount == 1 {
					currentSymbol.gearRatio = currentPart.partNumber
				}
				if currentSymbol.partCount == 2 {
					currentSymbol.gearRatio = currentSymbol.gearRatio * currentPart.partNumber
				}
				if currentSymbol.partCount > 2 {
					currentSymbol.gearRatio = 0
				}
			}
		}
	}
}

func processDay3A(inputLines []string) {
	fmt.Println("*** Day 3 part A ***")
	//total := 0

	//initialize objects needed
	var machineParts []machinePart
	var machineSymbols []machineSymbol

	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		getNumbersAndSymbols(inputLine, i, &machineParts, &machineSymbols)
		//i = len(inputLines)
	}

	fmt.Println("Number of machine parts: " + strconv.Itoa(len(machineParts)))
	fmt.Println("Number of symbols: " + strconv.Itoa(len(machineSymbols)))

	// Now that we have the listing of symbols and machine parts, let's see where a symbol is touching a part
	checkPartsTouchingSymbols(machineParts, machineSymbols)

	fmt.Println("*** End processing ***")
}

func processDay3B(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	//initialize objects needed
	var machineParts []machinePart
	var machineSymbols []machineSymbol

	// Read the input and separate out symbols and numbers
	fmt.Println("Input: ")
	for i := 0; i < len(inputLines); i++ {
		inputLine := inputLines[i]
		getNumbersAndSymbols(inputLine, i, &machineParts, &machineSymbols)
	}
	fmt.Println("Number of machine parts: " + strconv.Itoa(len(machineParts)))
	fmt.Println("Number of symbols: " + strconv.Itoa(len(machineSymbols)))

	// Now that we have the listing of symbols and machine parts,
	//  let's see where a symbol is touching a machine part
	checkPartsTouchingSymbols(machineParts, machineSymbols)

	// Calculate the ratio
	totalRatio := calculateRatio(machineSymbols)

	fmt.Println("Total for Ratio: " + strconv.Itoa(totalRatio))
	fmt.Println("*** End processing ***")
}

func calculateRatio(machineSymbols []machineSymbol) int {
	totalRatio := 0
	for s := 0; s < len(machineSymbols); s++ {
		if machineSymbols[s].partCount == 2 {
			//fmt.Println("Adding: " + strconv.Itoa(machineSymbols[s].gearRatio))
			totalRatio = totalRatio + machineSymbols[s].gearRatio
		}
	}
	return totalRatio
}
