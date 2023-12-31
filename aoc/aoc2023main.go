package main

import (
	"fmt"
	"strconv"

	"example.com/day"
	"example.com/utility"
)

func main() {
	aocDay := day.AdventOfCodeDay{
		10,
		"B",
		"10"}
	fmt.Println("*** Advent of Code 2023 ***")
	filename := getFilename(aocDay)

	linesRead := utility.ReadinputFile(filename)

	//fmt.Println("Read from file: " + strings.Join(linesRead, ""))

	switch strconv.Itoa(aocDay.Day) + aocDay.Part {
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

	switch aocDay.Day {
	case 4:
		//processDay4(linesRead, aocDay)
		break
	case 5:
		//processDay5(linesRead, aocDay)
		break

	case 6:
		day.ProcessDay6(linesRead, aocDay)
		break

	case 7:
		day.ProcessDay7(linesRead, aocDay)
		break

	case 8:
		day.ProcessDay8(linesRead, aocDay)
		break

	case 9:
		day.ProcessDay9(linesRead, aocDay)
		break

	case 10:
		day.ProcessDay10(linesRead, aocDay)
		break
	}

	fmt.Println("*** End of program ***")
}

func getFilename(aocDay day.AdventOfCodeDay) string {
	filename := "input" + strconv.Itoa(aocDay.Day) + aocDay.Part + ".txt"
	if len(aocDay.FileOverride) > 0 {
		filename = "input" + aocDay.FileOverride + ".txt"
	}
	return filename
}
