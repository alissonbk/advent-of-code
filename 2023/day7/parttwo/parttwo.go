package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var CARDS = [5]string{"T", "J", "Q", "K", "A"}

type Hand struct {
	hand string
	bid  uint
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
	var checkedRunes []rune
	var equals []string
	var numberEquals uint8 = 0
	runes := []rune(hand)
	numberOfJ := 0
	for _, r := range runes {
		if string(r) == "J" {
			numberOfJ += 1
		}
		for _, checkedRune := range checkedRunes {
			if r == checkedRune {
				alreadyInEquals := false
				if equals == nil {
					equals = append(equals, string(r))
				}
				for _, e := range equals {
					if string(r) == e {
						alreadyInEquals = true
					}
				}
				if !alreadyInEquals {
					equals = append(equals, string(r))
				}
				numberEquals += 1
			}
		}
		checkedRunes = append(checkedRunes, r)
	}
	// 10 because the loop will double the sum when it's all same value
	if numberEquals == 10 {
		return 7 // five of a kind
	}
	// 6 because the loop will redo sum once again
	if numberEquals == 6 {
		if len(equals) == 1 {
			if numberOfJ > 0 {
				return 7 // five of kind with J
			}
		}
	}
	if numberEquals == 4 {
		if len(equals) == 1 {
			if numberOfJ > 0 {
				return 7 // five of kind with J
			}
			return 6 // four of a kind
		}
		if len(equals) == 2 {
			if numberOfJ > 0 {
				return 7 // five of kind with J
			}
			return 5 // full house
		}
	}
	if numberEquals == 3 {
		if len(equals) == 1 {
			if numberOfJ > 0 {
				return 6 // four of a kind with J
			}
			return 4 // three of a kind
		}
	}
	if numberEquals == 2 {
		if len(equals) == 2 {
			if numberOfJ == 1 {
				return 5 // full house with J
			}
			if numberOfJ == 2 {
				return 6 // four of a kind with J
			}
			return 2 // two pair
		}
	}
	if numberEquals == 1 {
		if numberOfJ > 0 {
			return 4 // three of a king with J
		}
		return 1 // one pair
	}
	if numberEquals == 0 {
		if numberOfJ == 1 {
			return 1 // one pair
		}
	}
	return numberEquals
}

func isNumberOrJ(b byte) bool {
	return (b >= 48 && b <= 57) || b == 74
}

func getNumberOrJ(str string) int {
	var n int
	if str == "J" {
		n = 1
	} else {
		n, _ = strconv.Atoi(str)
	}

	return n
}

func appendHandsOrdered(hands *[]Hand, hand Hand) {
	handsLen := len(*hands)
	newHandStr := []rune(hand.hand)
	if handsLen == 0 {
		*hands = append(*hands, hand)
		return
	}

	for i := 0; i < handsLen; i++ {
		handStr := []rune((*hands)[i].hand)

		idx := 0
		for {
			oldHand := handStr[idx]
			newHand := newHandStr[idx]
			for oldHand == newHand {
				if idx+1 < len(handStr) {
					idx++
					oldHand = handStr[idx]
					newHand = newHandStr[idx]
				} else {
					utils.InsertAtIndex(hands, i, hand)
					return
				}

			}

			if isNumberOrJ(byte(oldHand)) && isNumberOrJ(byte(newHand)) {
				oldN := getNumberOrJ(string(oldHand))
				newN := getNumberOrJ(string(newHand))
				if newN < oldN {
					utils.InsertAtIndex(hands, i, hand)
					return
				} else {
					if i == len(*hands)-1 {
						*hands = append(*hands, hand)
						return
					}
					break
				}
			} else if !isNumberOrJ(byte(oldHand)) && isNumberOrJ(byte(newHand)) {
				utils.InsertAtIndex(hands, i, hand)
				return
			} else if isNumberOrJ(byte(oldHand)) && !isNumberOrJ(byte(newHand)) {
				if i == len(*hands)-1 {
					*hands = append(*hands, hand)
					return
				}
				break
			} else if getCardStrength(string(oldHand)) > getCardStrength(string(newHand)) {
				utils.InsertAtIndex(hands, i, hand)
				return
			} else {
				if i == handsLen-1 {
					*hands = append(*hands, hand)
					return
				}
				break
			}

		}
	}
}

func Run() {
	rows := utils.GetSliceFromFile("/day7/parttwo/input.txt")
	handMap := make(map[uint8][]Hand)
	totalSum := 0
	for _, row := range rows {
		split := strings.Split(row, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])
		handType := getHandType(hand)
		handsByType := handMap[handType]
		appendHandsOrdered(&handsByType, Hand{hand: hand, bid: uint(bid)})
		handMap[handType] = handsByType
	}

	counter := uint(1)
	i := uint8(0)

	for i <= 7 {
		hands := handMap[i]
		if hands != nil {
			if len(hands) == 1 {
				totalSum += int(hands[0].bid * counter)
				counter++
			} else {
				for _, hand := range hands {
					totalSum += int(hand.bid * counter)
					counter++
				}
			}
		}
		i++
	}
	fmt.Println(totalSum)
}
