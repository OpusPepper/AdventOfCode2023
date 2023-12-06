package utility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadinputFile(aocDay string) []string {
	fmt.Println("*** ReadinputFile ***")

	fileName := "../input/" + aocDay
	fmt.Println("Reading file: " + fileName)

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	check(err)

	fmt.Println("lines length: " + strconv.Itoa(len(lines)))
	fmt.Println("*** End ReadinputFile ***")
	return lines
}
