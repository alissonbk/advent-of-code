package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func appendOnlyValidGameID(allGameSets []gameSet, validGamesIDSum *int, gameID int) {
	for _, gs := range allGameSets {
		for _, gp := range gs {
			if gp.valid == false {
				return
			}
		}
	}
	*validGamesIDSum += gameID
}

func Run() {
	rows := utils.GetSliceFromFile("/day2/partone/input.txt")
	validGamesIDSum := 0
	for _, row := range rows {
		splitted := strings.Split(row, ": ")
		gameIDStr, _ := strings.CutPrefix(splitted[0], "Game ")
		gameIDStr = strings.Replace(gameIDStr, ":", "", 1)
		gameID, _ := strconv.Atoi(gameIDStr)

		gameSetsString := strings.Split(splitted[1], ";")
		var allGameSets []gameSet
		for _, str := range gameSetsString {
			separated := strings.Split(str, ",")
			var gameSet []gamePart
			for _, sp := range separated {
				sp = strings.Trim(sp, " ")
				s := strings.Split(sp, " ")
				qty, _ := strconv.Atoi(s[0])
				gp := gamePart{
					color: s[1],
					qty:   qty,
					valid: true,
				}
				gp.validate()
				gameSet = append(gameSet, gp)
			}
			allGameSets = append(allGameSets, gameSet)
		}
		appendOnlyValidGameID(allGameSets, &validGamesIDSum, gameID)

	}
	fmt.Println("Part one result: ", validGamesIDSum)
}
