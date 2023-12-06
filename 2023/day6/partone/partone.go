package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func getSliceFromRow(row string) []int {
	var slice []int
	split := strings.Split(row, " ")
	for _, s := range split {
		if s != "" {
			n, _ := strconv.Atoi(s)
			slice = append(slice, n)
		}
	}
	return slice
}

func Run() {
	rows := utils.GetSliceFromFile("/day6/partone/input.txt")

	timeRow, _ := strings.CutPrefix(rows[0], "Time:")
	timeRow = strings.Trim(timeRow, " ")
	timeSlice := getSliceFromRow(timeRow)
	distanceRow, _ := strings.CutPrefix(rows[1], "Distance:")
	distanceRow = strings.Trim(distanceRow, " ")
	distanceSlice := getSliceFromRow(distanceRow)

	totalWaysToWin := 1

	for idx := 0; idx < len(timeSlice); idx++ {
		var waysToWin int
		for i := 1; i < timeSlice[idx]-1; i++ {
			timeRemaining := timeSlice[idx] - i
			totalDistance := timeRemaining * i
			if totalDistance > distanceSlice[idx] {
				waysToWin += 1
			}
		}
		if waysToWin > 0 {
			totalWaysToWin *= waysToWin
		}
	}

	fmt.Println(totalWaysToWin)
}
