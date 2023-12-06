package main

import (
	"fmt"
	"strconv"
)

type AdventOfCodeDay struct {
	day          int
	part         string
	fileOverride string
}

func main() {
	aocDay := AdventOfCodeDay{
		day:          6,
		part:         "B",
		fileOverride: "6"}

	fmt.Println("*** Advent of Code 2023 ***")
	filename := getFilename(aocDay)

	linesRead := ReadinputFile(filename)

	//fmt.Println("Read from file: " + strings.Join(linesRead, ""))

	switch strconv.Itoa(aocDay.day) + aocDay.part {
	case "1A":
		//processDay1A(linesRead)
		break
	case "1B":
		//processDay1B(linesRead)
		break
	case "2A":
		//processDay2A(linesRead)
		break
	case "2B":
		//processDay2B(linesRead)
		break
	case "3A":
		//processDay3A(linesRead)
		break
	case "3B":
		//processDay3B(linesRead, aocDay)
		break
	}

	switch aocDay.day {
	case 4:
		//processDay4(linesRead, aocDay)
		break
	case 5:
		//processDay5(linesRead, aocDay)
		break

	case 6:
		processDay6(linesRead, aocDay)
		break
	}

	fmt.Println("*** End of program ***")
}

func getFilename(aocDay AdventOfCodeDay) string {
	filename := "input" + strconv.Itoa(aocDay.day) + aocDay.part + ".txt"
	if len(aocDay.fileOverride) > 0 {
		filename = "input" + aocDay.fileOverride + ".txt"
	}
	return filename
}
