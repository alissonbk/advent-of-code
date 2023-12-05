package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	rows := utils.GetSliceFromFile("/day4/partone/input.txt")
	totalSum := 0
	for idx, row := range rows {
		rowSum := 0
		foundFirst := false
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
					if !foundFirst {
						foundFirst = true
						rowSum += 1
					} else {
						rowSum += rowSum
					}
				}
			}
		}
		totalSum += rowSum
	}
	fmt.Println(totalSum)
}
