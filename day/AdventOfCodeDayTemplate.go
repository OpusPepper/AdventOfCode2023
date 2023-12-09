package day

import (
	"fmt"
	"strconv"
)

func ProcessDayX(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	if aocDay.Part == "A" {
		processDayXA(inputLines)
	} else {
		processDayXB(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDayXA(inputLines []string) {
	fmt.Println("*** Start part A ***")

	fmt.Println("*** End part A ***")
}

func processDayXB(inputLines []string) {
	fmt.Println("*** Start part B ***")

	fmt.Printf("*** End part B ***")
}
