package main

import (
	"fmt"
)

func main() {
	var aocDay = "3A"

	fmt.Println("*** Advent of Code 2023 ***")

	linesRead := ReadinputFile(string(aocDay))

	//fmt.Println("Read from file: " + strings.Join(linesRead, ""))

	processDay3B(linesRead)

	fmt.Println("*** End of program ***")
}
