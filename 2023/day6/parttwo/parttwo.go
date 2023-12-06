package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func getSliceFromRow(row string) int {
	numberAsStr := strings.ReplaceAll(row, " ", "")
	n, _ := strconv.Atoi(numberAsStr)
	return n
}

func Run() {
	rows := utils.GetSliceFromFile("/day6/parttwo/input.txt")

	timeRow, _ := strings.CutPrefix(rows[0], "Time:")
	timeRow = strings.Trim(timeRow, " ")
	raceTime := getSliceFromRow(timeRow)
	distanceRow, _ := strings.CutPrefix(rows[1], "Distance:")
	distanceRow = strings.Trim(distanceRow, " ")
	raceDistance := getSliceFromRow(distanceRow)

	var waysToWin int
	for i := 1; i < raceTime-1; i++ {
		timeRemaining := raceTime - i
		totalDistance := timeRemaining * i
		if totalDistance > raceDistance {
			waysToWin += 1
		}
	}

	fmt.Println(waysToWin)
}
