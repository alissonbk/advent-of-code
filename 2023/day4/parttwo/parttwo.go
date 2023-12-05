package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	id           int
	winnerPoints int
	instances    int
}

type RepeatedCard struct {
	id            int
	repeatedTimes int
}
type RepeatedCards []RepeatedCard

func (rcds *RepeatedCards) getById(id int) *RepeatedCard {
	for _, rc := range *rcds {
		if rc.id == id {
			return &rc
		}
	}
	return nil
}

func (rcds *RepeatedCards) getRepeatedInstancesById(id int) int {
	for _, rc := range *rcds {
		if rc.id == id {
			return rc.repeatedTimes
		}
	}
	return 0
}

func (rcds *RepeatedCards) increaseRepeatedTimesById(id int) {
	for idx, rc := range *rcds {
		if rc.id == id {
			rc.repeatedTimes += 1
			(*rcds)[idx] = rc
		}
	}
}

func Run() {
	rows := utils.GetSliceFromFile("/day4/parttwo/input.txt")
	totalSum := 0
	var allCards []Card
	var repeatedCards RepeatedCards
	for idx, row := range rows {
		var card Card
		card.id = idx + 1
		card.instances = repeatedCards.getRepeatedInstancesById(card.id) + 1 // 1 original
		cardSum := 0
		var winnerTickets []int
		var myTickets []int

		numberAsString := ""
		cardName := "Card " + strconv.Itoa(idx+1) + ": "
		row = strings.Trim(strings.Replace(row, cardName, "", -1), "")
		sepIndex := strings.Index(row, "|")
		chars := []rune(row)

		for i := 0; i < sepIndex; i++ {
			charAsString := string(chars[i])
			if charAsString != " " {
				numberAsString += charAsString
			} else {
				number, _ := strconv.Atoi(numberAsString)
				winnerTickets = append(winnerTickets, number)
				numberAsString = ""
			}
		}
		for i := sepIndex + 1; i < len(row); i++ {
			charAsString := string(chars[i])
			if charAsString != " " {
				numberAsString += charAsString
			} else if numberAsString != "" {
				number, _ := strconv.Atoi(numberAsString)
				myTickets = append(myTickets, number)
				numberAsString = ""
			}
			if charAsString != "" && i == len(row)-1 {
				number, _ := strconv.Atoi(numberAsString)
				myTickets = append(myTickets, number)
				numberAsString = ""
			}
		}
		for _, ticket := range myTickets {
			for _, wTicket := range winnerTickets {
				if ticket == wTicket {
					cardSum += 1
				}
			}
		}
		card.winnerPoints = cardSum //maybe useless
		allCards = append(allCards, card)

		for idx := 0; idx < card.instances; idx++ {
			for i := 1; i <= card.winnerPoints; i++ {
				rc := repeatedCards.getById(card.id + i)
				if rc == nil {
					rc = &RepeatedCard{
						id:            card.id + i,
						repeatedTimes: 1}
					repeatedCards = append(repeatedCards, *rc)
				} else {
					repeatedCards.increaseRepeatedTimesById(rc.id)
				}

			}
		}
	}
	for _, c := range allCards {
		totalSum += c.instances
	}
	fmt.Println(totalSum)
}
