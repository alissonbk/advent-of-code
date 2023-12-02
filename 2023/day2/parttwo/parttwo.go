package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func sumPowerOfMinimunValues(allGameSets []gameSet, finalPowerSumOfGames *int) {
	minGreen := 1
	minRed := 1
	minBlue := 1
	for _, gs := range allGameSets {
		for _, gp := range gs {
			switch gp.color {
			case "green":
				if gp.qty > minGreen {
					minGreen = gp.qty
				}
				break
			case "red":
				if gp.qty > minRed {
					minRed = gp.qty
				}
				break
			case "blue":
				if gp.qty > minBlue {
					minBlue = gp.qty
				}
				break
			default:
				panic("Color not found!")
			}
		}
	}
	*finalPowerSumOfGames += minRed * minBlue * minGreen
}

func Run() {
	rows := utils.GetSliceFromFile("/day2/parttwo/input.txt")
	finalPowerSumOfGames := 0
	for _, row := range rows {
		splitted := strings.Split(row, ": ")
		gameSetsString := strings.Split(splitted[1], ";")
		var allGameSets []gameSet
		for _, str := range gameSetsString { // one set
			separated := strings.Split(str, ",")
			var gameSet []gamePart
			for _, sp := range separated { // one color/number
				sp = strings.Trim(sp, " ")
				s := strings.Split(sp, " ")
				qty, _ := strconv.Atoi(s[0])
				gp := gamePart{
					color: s[1],
					qty:   qty,
				}
				gameSet = append(gameSet, gp)
			}
			allGameSets = append(allGameSets, gameSet)
		}
		sumPowerOfMinimunValues(allGameSets, &finalPowerSumOfGames)

	}
	fmt.Println("Part two result: ", finalPowerSumOfGames)
}
