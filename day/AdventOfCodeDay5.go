package day

import (
	"fmt"
	"strconv"
	"strings"
)

type aMap struct {
	position         int
	typeOfMap        string
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func getSeeds(inputLine string) []int {
	seedNumbers := strings.Fields(strings.Split(inputLine, ":")[1])
	var seedNumbersInt []int
	for i := 0; i < len(seedNumbers); i++ {
		num, err := strconv.Atoi(seedNumbers[i])
		check(err)
		seedNumbersInt = append(seedNumbersInt, num)
	}
	return seedNumbersInt
}

func processDay5(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	//initialize objects needed

	if aocDay.Part == "A" {
		processDay5A(inputLines)
	} else {
		processDay5B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func getSeedOutput(mapper []aMap, seedNum int, mapToUse int) int {
	destinationVal := -1
	foundDestinationValue := false
	for i := 0; i < len(mapper); i++ {
		currentMap := mapper[i]
		if mapper[i].position == mapToUse {
			if currentMap.sourceStart <= seedNum && seedNum <= (currentMap.sourceStart+(currentMap.rangeLength-1)) {
				// we found a match in source, now map to destination
				//fmt.Println(strconv.Itoa(currentMap.sourceStart) + " <= " + strconv.Itoa(seedNum) + " <= " + strconv.Itoa((currentMap.sourceStart + (currentMap.rangeLength - 1))))
				destinationVal = currentMap.destinationStart + (seedNum - currentMap.sourceStart)
				foundDestinationValue = true
				continue
			}
		}
	}

	if foundDestinationValue {
		return destinationVal
	}

	return seedNum
}

func processDay5A(inputLines []string) {
	fmt.Println("Input: ")
	var mapper []aMap
	var seeds []int
	var currentMap aMap = aMap{0, "", 0, 0, 0}
	mapCounter := 0
	mapName := ""
	if currentMap != (aMap{}) {
		fmt.Println("Not nil!")
	}

	for i := 0; i < len(inputLines); i++ {

		inputLine := inputLines[i]
		if len(inputLine) == 0 {
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "seeds") {
			seeds = getSeeds(inputLine)
			fmt.Println(seeds)
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "map") {
			mapName = strings.Fields(inputLine)[0]
			mapCounter = mapCounter + 1
			continue
		}
		destSourceRange := strings.Fields(inputLine)
		destStart, _ := strconv.Atoi(destSourceRange[0])
		sourceStart, _ := strconv.Atoi(destSourceRange[1])
		rangeVal, _ := strconv.Atoi(destSourceRange[2])
		currentMap = aMap{
			mapCounter,
			mapName,
			destStart,
			sourceStart,
			rangeVal}
		mapper = append(mapper, currentMap)
	}

	if len(mapper) != 0 {
		totalValStr := strconv.Itoa(len(mapper))
		fmt.Println("We have some mappers! " + totalValStr)
	}
	printOutMaps(mapper)
	if len(seeds) > 0 {
		totalSeedsValStr := strconv.Itoa(len(seeds))
		fmt.Println("We have some seeds! " + totalSeedsValStr)
	}

	// Let's work a seed through the steps now
	lowestNumber := -1
	for i := 0; i < len(seeds); i++ {
		//fmt.Println("Looking at seed: " + strconv.Itoa(seeds[i]))
		seedToSoilOutput := getSeedOutput(mapper, seeds[i], 1)
		//fmt.Println(" After map 1: " + strconv.Itoa(seedToSoilOutput))
		soilToFertilizerOutput := getSeedOutput(mapper, seedToSoilOutput, 2)
		//fmt.Println(" After map 2: " + strconv.Itoa(soilToFertilizerOutput))
		fertilizerToWaterOutput := getSeedOutput(mapper, soilToFertilizerOutput, 3)
		//fmt.Println(" After map 3: " + strconv.Itoa(fertilizerToWaterOutput))
		waterToLightOutput := getSeedOutput(mapper, fertilizerToWaterOutput, 4)
		//fmt.Println(" After map 4: " + strconv.Itoa(waterToLightOutput))
		lightToTemperatureOutput := getSeedOutput(mapper, waterToLightOutput, 5)
		//fmt.Println(" After map 5: " + strconv.Itoa(lightToTemperatureOutput))
		temperatureToHumidityOutput := getSeedOutput(mapper, lightToTemperatureOutput, 6)
		//fmt.Println(" After map 6: " + strconv.Itoa(temperatureToHumidityOutput))
		humidityToLocationOutput := getSeedOutput(mapper, temperatureToHumidityOutput, 7)
		//fmt.Println(" After map 7: " + strconv.Itoa(humidityToLocationOutput))

		if i == 0 {
			lowestNumber = humidityToLocationOutput
			continue
		}
		if humidityToLocationOutput < lowestNumber {
			lowestNumber = humidityToLocationOutput
		}
	}

	fmt.Println("lowestNumber: " + strconv.Itoa(lowestNumber))
}

func processDay5B(inputLines []string) {
	fmt.Println("Input: ")
	var mapper []aMap
	var seeds []int
	var currentMap aMap = aMap{0, "", 0, 0, 0}
	mapCounter := 0
	mapName := ""
	if currentMap != (aMap{}) {
		fmt.Println("Not nil!")
	}

	for i := 0; i < len(inputLines); i++ {

		inputLine := inputLines[i]
		if len(inputLine) == 0 {
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "seeds") {
			seeds = getSeeds(inputLine)
			fmt.Println(seeds)
			continue
		}
		if strings.Contains(strings.ToLower(inputLine), "map") {
			mapName = strings.Fields(inputLine)[0]
			mapCounter = mapCounter + 1
			continue
		}
		destSourceRange := strings.Fields(inputLine)
		destStart, _ := strconv.Atoi(destSourceRange[0])
		sourceStart, _ := strconv.Atoi(destSourceRange[1])
		rangeVal, _ := strconv.Atoi(destSourceRange[2])
		currentMap = aMap{
			mapCounter,
			mapName,
			destStart,
			sourceStart,
			rangeVal}
		mapper = append(mapper, currentMap)
	}

	if len(mapper) != 0 {
		totalValStr := strconv.Itoa(len(mapper))
		fmt.Println("We have some mappers! " + totalValStr)
	}
	printOutMaps(mapper)
	if len(seeds) > 0 {
		totalSeedsValStr := strconv.Itoa(len(seeds))
		fmt.Println("We have some seeds! " + totalSeedsValStr)
	}

	// Let's work a seed through the steps now
	lowestNumber := -1
	for i := 0; i < len(seeds); i = i + 2 {
		fmt.Println("Processing seed pair: " + strconv.Itoa(seeds[i]) + ", " + strconv.Itoa(seeds[i+1]))
		for j := 0; j < (seeds[i+1]); j++ {
			currentSeed := seeds[i] + j

			//fmt.Println("Looking at seed: " + strconv.Itoa(seeds[i]))
			seedToSoilOutput := getSeedOutput(mapper, currentSeed, 1)
			//fmt.Println(" After map 1: " + strconv.Itoa(seedToSoilOutput))
			soilToFertilizerOutput := getSeedOutput(mapper, seedToSoilOutput, 2)
			//fmt.Println(" After map 2: " + strconv.Itoa(soilToFertilizerOutput))
			fertilizerToWaterOutput := getSeedOutput(mapper, soilToFertilizerOutput, 3)
			//fmt.Println(" After map 3: " + strconv.Itoa(fertilizerToWaterOutput))
			waterToLightOutput := getSeedOutput(mapper, fertilizerToWaterOutput, 4)
			//fmt.Println(" After map 4: " + strconv.Itoa(waterToLightOutput))
			lightToTemperatureOutput := getSeedOutput(mapper, waterToLightOutput, 5)
			//fmt.Println(" After map 5: " + strconv.Itoa(lightToTemperatureOutput))
			temperatureToHumidityOutput := getSeedOutput(mapper, lightToTemperatureOutput, 6)
			//fmt.Println(" After map 6: " + strconv.Itoa(temperatureToHumidityOutput))
			humidityToLocationOutput := getSeedOutput(mapper, temperatureToHumidityOutput, 7)
			//fmt.Println(" After map 7: " + strconv.Itoa(humidityToLocationOutput))
			//fmt.Println("Seed: s(" + strconv.Itoa(currentSeed) +
			//		"), " + strconv.Itoa(seedToSoilOutput) +
			//		", " + strconv.Itoa(soilToFertilizerOutput) +
			//		", " + strconv.Itoa(fertilizerToWaterOutput) +
			//		", " + strconv.Itoa(waterToLightOutput) +
			//		", " + strconv.Itoa(lightToTemperatureOutput) +
			//		", " + strconv.Itoa(temperatureToHumidityOutput) +
			//		", " + strconv.Itoa(humidityToLocationOutput))

			if i == 0 && j == 0 {
				lowestNumber = humidityToLocationOutput
				continue
			}
			if humidityToLocationOutput < lowestNumber {
				lowestNumber = humidityToLocationOutput
			}

		}

	}

	fmt.Println("lowestNumber: " + strconv.Itoa(lowestNumber))
}

func printOutMaps(mapper []aMap) {
	for i := 0; i < 7; i++ {
		for m := 0; m < len(mapper); m++ {
			if mapper[m].position == (i + 1) {
				fmt.Println(strconv.Itoa(mapper[m].position) + ": " + strconv.Itoa(mapper[m].destinationStart) + ", " + strconv.Itoa(mapper[m].sourceStart) + ", " + strconv.Itoa(mapper[m].rangeLength))
			}
		}
	}
}
