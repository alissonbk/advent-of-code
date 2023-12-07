package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// const (
//
//	A = iota
//	K
//	Q
//	J
//	T
//
// )
var CARDS = [5]string{"T", "J", "Q", "K", "A"}

type Hand struct {
	hand string
	bid  uint16
}

func getCardStrength(card string) int {
	for idx, c := range CARDS {
		if card == c {
			return idx
		}
	}
	panic("Strength not found")
}

func getHandType(hand string) uint8 {
	runes := []rune(hand)
	var checkedRunes []rune
	var numberEquals uint8 = 0
	for _, r := range runes {
		for _, checkedRune := range checkedRunes {
			if r == checkedRune {
				numberEquals += 1
			}
		}
		checkedRunes = append(checkedRunes, r)
	}
	return numberEquals
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func appendHandsOrdered(hands *[]Hand, hand Hand) {
	handsLen := len(*hands)
	//var newHands []Hand
	newHandStr := []rune(hand.hand)
	//posToAppend := 0
	//foundFinalPosition := false
	if handsLen == 0 {
		*hands = append(*hands, hand)
		return
	}

	for i := 0; i < handsLen; i++ {
		handStr := []rune((*hands)[i].hand)

		for idx := 0; idx < len(handStr); idx++ {
			oldHand := handStr[idx]
			newHand := newHandStr[idx]
			for oldHand == newHand {
				idx++
				oldHand = handStr[idx]
				newHand = newHandStr[idx]
			}

			if isNumber(byte(oldHand)) && isNumber(byte(newHand)) {
				oldN, _ := strconv.Atoi(string(oldHand))
				newN, _ := strconv.Atoi(string(newHand))
				if newN < oldN {
					utils.InsertAtIndex(hands, i, hand)
					return
				} else {
					//posToAppend = i + 1
				}
			} else if !isNumber(byte(oldHand)) && isNumber(byte(newHand)) {
				utils.InsertAtIndex(hands, i, hand)
				return
			} else if isNumber(byte(oldHand)) && !isNumber(byte(newHand)) {
				//posToAppend = i + 1
			} else if getCardStrength(string(oldHand)) > getCardStrength(string(newHand)) {
				utils.InsertAtIndex(hands, i, hand)
				return
			} else {
				//posToAppend = i + 1
			}

		}
	}
}

func Run() {
	rows := utils.GetSliceFromFile("/day7/partone/input.txt")
	handMap := make(map[uint8][]Hand)
	counter := uint16(1)
	totalSum := 0
	for _, row := range rows {
		split := strings.Split(row, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])
		handType := getHandType(hand)
		handsByType := handMap[handType]
		appendHandsOrdered(&handsByType, Hand{hand: hand, bid: uint16(bid)})
		//handsByType = append(handsByType, Hand{hand, uint16(bid)})
		handMap[getHandType(hand)] = handsByType
	}

	fmt.Println(handMap)
	for i := uint8(0); i < uint8(len(handMap)); i++ {
		hands := handMap[i]
		if hands != nil {
			if len(hands) == 1 {
				totalSum += int(hands[0].bid * counter)
				counter++
			} else {
				// should be ordered
				//for _, hand := range hands {
				//	var ordered []uint16
				//
				//}
			}

		}

	}

	fmt.Println(totalSum)
}
