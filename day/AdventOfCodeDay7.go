package day

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type cardHand struct {
	originalString string
	bid            int
	rank           int
	strength       int
}

type myCard struct {
	cardValue              rune
	numberOfTimesItAppears int
}

const (
	fiveOfAKind  int = 7
	fourOfAKind  int = 6
	fullHouse    int = 5
	threeOfAKind int = 4
	twoPair      int = 3
	onePair      int = 2
	highCard     int = 1
	nothing      int = 0
)

const (
	ace   int = 14
	king  int = 13
	queen int = 12
	jack  int = 0
	ten   int = 10
)

func evaluateHand(inCardHand cardHand) int {
	var totalCards []myCard = make([]myCard, 0)
	for _, c := range inCardHand.originalString {
		foundCard := false
		for i := 0; i < len(totalCards); i++ {

			if totalCards[i].cardValue == c {
				totalCards[i].numberOfTimesItAppears++
				foundCard = true
				break
			}
		}
		if !foundCard {
			newCard := myCard{
				c,
				1}
			totalCards = append(totalCards, newCard)
		}
	}
	fmt.Println("totalCards: " + strconv.Itoa(len(totalCards)))

	if len(totalCards) == 1 { // five of a kind
		return fiveOfAKind
	}
	if len(totalCards) == 2 { // four of a kind or full house
		if totalCards[0].numberOfTimesItAppears == 2 || totalCards[0].numberOfTimesItAppears == 3 {
			return fullHouse
		}
		return fourOfAKind
	}
	if len(totalCards) == 3 { // three of a kind or two pair
		if totalCards[0].numberOfTimesItAppears == 3 || totalCards[1].numberOfTimesItAppears == 3 || totalCards[2].numberOfTimesItAppears == 3 {
			return threeOfAKind
		}
		return twoPair
	}
	if len(totalCards) == 4 { // one pair
		return onePair
	}
	return highCard
}

func ProcessDay7(inputLines []string, aocDay AdventOfCodeDay) {
	fmt.Println("*** Day " + strconv.Itoa(aocDay.Day) + " part " + aocDay.Part + " ***")

	//initialize objects needed

	if aocDay.Part == "A" {
		processDay7A(inputLines)
	} else {
		processDay7B(inputLines)
	}

	fmt.Println("*** End processing ***")
}

func processDay7A(inputLines []string) {
	fmt.Println("Input: ")
	var fullHand []cardHand = make([]cardHand, 0)
	for i := 0; i < len(inputLines); i++ {
		tempBid, _ := strconv.Atoi(strings.Fields(inputLines[i])[1])
		newCardHand := cardHand{
			originalString: strings.Fields(inputLines[i])[0],
			bid:            tempBid,
			rank:           0,
			strength:       nothing}
		findStrength := evaluateHand(newCardHand)
		newCardHand.strength = findStrength
		fullHand = append(fullHand, newCardHand)

		fmt.Println("Card: " + newCardHand.originalString + ", bid: " + strconv.Itoa(newCardHand.bid) + ", rank: " + strconv.Itoa(newCardHand.rank) + ", strength: " + strconv.Itoa(newCardHand.strength))
	}

	By(func(i, j *cardHand) bool {
		if i.strength > j.strength {
			return true
		}
		if i.strength < j.strength {
			return false
		}

		for k := 0; k < len(i.originalString); k++ {
			charI := string((i.originalString)[k])
			charJ := string((j.originalString)[k])
			if charI == charJ {
				continue
			}
			if charToConst(charI) > charToConst(charJ) {
				return true
			}
			return false
		}
		return false
	}).Sort(fullHand)

	// Add rank
	currentRank := len(fullHand)
	for r := 0; r < len(fullHand); r++ {
		fullHand[r].rank = currentRank
		currentRank--
	}

	printCardHands(fullHand)

	// Calculate total bids
	total := 0
	for _, t := range fullHand {
		handWinnings := t.bid * t.rank
		total += handWinnings
		fmt.Println("Winnings: " + strconv.Itoa(handWinnings))
	}
	fmt.Println("grandTotal: " + strconv.Itoa(total))
}

func printCardHands(fullHand []cardHand) {
	fmt.Println("Printing fullhand: ")
	for i := 0; i < len(fullHand); i++ {
		fmt.Println("Card: " + fullHand[i].originalString + ", bid: " + strconv.Itoa(fullHand[i].bid) + ", rank: " + strconv.Itoa(fullHand[i].rank) + ", strength: " + strconv.Itoa(fullHand[i].strength))
	}
}

type By func(p1, p2 *cardHand) bool

func (by By) Sort(cardHards []cardHand) {
	ps := &fullHandSorter{
		cardHands: cardHards,
		by:        by,
	}
	sort.Sort(ps)
}

type fullHandSorter struct {
	cardHands []cardHand
	by        func(p1, p2 *cardHand) bool
}

func (s *fullHandSorter) Len() int {
	return len(s.cardHands)
}

func (s *fullHandSorter) Swap(i, j int) {
	s.cardHands[i], s.cardHands[j] = s.cardHands[j], s.cardHands[i]
}

func (s *fullHandSorter) Less(i, j int) bool {
	return s.by(&s.cardHands[i], &s.cardHands[j])
}

func compareHand(i cardHand, j cardHand) bool {
	for k := 0; k < len(i.originalString); k++ {
		charI := string((i.originalString)[k])
		charJ := string((j.originalString)[k])
		if charI == charJ {
			continue
		}
		if charToConst(charI) > charToConst(charJ) {
			return true
		}
		return false
	}
	return false
}

func charToConst(i string) int {
	if i == "A" {
		return ace
	}
	if i == "K" {
		return king
	}
	if i == "Q" {
		return queen
	}
	if i == "J" {
		return jack
	}
	if i == "T" {
		return ten
	}
	val, _ := strconv.Atoi(i)
	return val
}

func processDay7B(inputLines []string) {
	fmt.Println("Input: ")
	var fullHand []cardHand = make([]cardHand, 0)
	for i := 0; i < len(inputLines); i++ {
		tempBid, _ := strconv.Atoi(strings.Fields(inputLines[i])[1])
		newCardHand := cardHand{
			originalString: strings.Fields(inputLines[i])[0],
			bid:            tempBid,
			rank:           0,
			strength:       nothing}
		findStrength := evaluateHandWithJokers(newCardHand)
		newCardHand.strength = findStrength
		fullHand = append(fullHand, newCardHand)

		//fmt.Println("Card: " + newCardHand.originalString + ", bid: " + strconv.Itoa(newCardHand.bid) + ", rank: " + strconv.Itoa(newCardHand.rank) + ", strength: " + strconv.Itoa(newCardHand.strength))
	}

	By(func(i, j *cardHand) bool {
		if i.strength > j.strength {
			return true
		}
		if i.strength < j.strength {
			return false
		}

		for k := 0; k < len(i.originalString); k++ {
			charI := string((i.originalString)[k])
			charJ := string((j.originalString)[k])
			if charI == charJ {
				continue
			}
			if charToConst(charI) > charToConst(charJ) {
				return true
			}
			return false
		}
		return false
	}).Sort(fullHand)

	// Add rank
	currentRank := len(fullHand)
	for r := 0; r < len(fullHand); r++ {
		fullHand[r].rank = currentRank
		currentRank--
	}

	printCardHands(fullHand)

	// Calculate total bids
	total := 0
	for _, t := range fullHand {
		handWinnings := t.bid * t.rank
		total += handWinnings
		fmt.Println("Winnings: " + strconv.Itoa(handWinnings))
	}
	fmt.Println("grandTotal: " + strconv.Itoa(total))
}

func evaluateHandWithJokers(inCardHand cardHand) int {
	var totalCards []myCard = make([]myCard, 0)
	for _, c := range inCardHand.originalString {
		foundCard := false
		for i := 0; i < len(totalCards); i++ {

			if totalCards[i].cardValue == c {
				totalCards[i].numberOfTimesItAppears++
				foundCard = true
				break
			}
		}
		if !foundCard {
			newCard := myCard{
				c,
				1}
			totalCards = append(totalCards, newCard)
		}
	}
	//fmt.Println("totalCards: " + strconv.Itoa(len(totalCards)))

	hasJacks := doesHandHaveJacks(totalCards)
	numberOfJacks := 0
	for _, j := range totalCards {
		if j.cardValue == []rune("J")[0] {
			numberOfJacks = j.numberOfTimesItAppears
		}
	}

	if len(totalCards) == 1 { // five of a kind
		return fiveOfAKind
	}
	if len(totalCards) == 2 { // four of a kind or full house
		if totalCards[0].numberOfTimesItAppears == 2 || totalCards[0].numberOfTimesItAppears == 3 {
			if hasJacks {
				if numberOfJacks == 1 {
					return fourOfAKind
				}
				if numberOfJacks == 2 || numberOfJacks == 3 {
					return fiveOfAKind
				}
			}
			return fullHouse
		}
		if hasJacks {
			return fiveOfAKind
		}
		return fourOfAKind
	}
	if len(totalCards) == 3 { // three of a kind or two pair
		if totalCards[0].numberOfTimesItAppears == 3 || totalCards[1].numberOfTimesItAppears == 3 || totalCards[2].numberOfTimesItAppears == 3 {
			if hasJacks {
				if numberOfJacks == 1 || numberOfJacks == 3 {
					return fourOfAKind
				}
				if numberOfJacks == 2 {
					return fiveOfAKind
				}
			}
			return threeOfAKind
		}
		if hasJacks {
			if numberOfJacks == 1 {
				return fullHouse
			}
			if numberOfJacks == 2 {
				return fourOfAKind
			}
		}
		return twoPair
	}
	if len(totalCards) == 4 { // one pair
		if hasJacks {
			if numberOfJacks == 1 {
				return threeOfAKind
			}
			if numberOfJacks == 2 {
				return fourOfAKind
			}
			if numberOfJacks == 3 {
				return fiveOfAKind
			}
		}
		return onePair
	}
	return highCard
}

func doesHandHaveJacks(myCards []myCard) bool {
	foundJacks := false
	for _, j := range myCards {
		if j.cardValue == []rune("J")[0] {
			foundJacks = true
		}
	}
	return foundJacks
}
