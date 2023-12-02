package main

import (
	"fmt"
)

func main() {
	var aocDay = "2A"

	fmt.Println("*** Advent of Code 2023 ***")

	linesRead := ReadinputFile(string(aocDay))

	//fmt.Println("Read from file: " + strings.Join(linesRead, ""))

	processDay2B(linesRead)

	fmt.Println("*** End of program ***")
}
