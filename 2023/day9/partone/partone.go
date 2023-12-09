package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func getNewDifferenceLine(row []int) ([]int, bool) {
	var newS []int
	sum := 0
	allZeros := false
	for i := 0; i < len(row)-1; i++ {
		newS = append(newS, row[i+1]-row[i])
		sum += newS[i]
	}

	difThanZero := false
	for i := 0; i < len(newS); i++ {
		if newS[i] != 0 {
			difThanZero = true
			break
		}
	}
	allZeros = !difThanZero

	return newS, allZeros
}

func getHistoryPrediction(row []int) int {
	var rowStack [][]int
	allZeros := false

	for allZeros == false {
		rowStack = append(rowStack, row)
		r, az := getNewDifferenceLine(row)
		allZeros = az
		row = r
	}
	sum := 0
	for i := len(rowStack) - 1; i >= 0; i-- {
		sum += rowStack[i][len(rowStack[i])-1]
	}
	return sum
}

func Run() {
	rows := utils.GetSliceFromFile("/day9/partone/input.txt")
	totalSum := 0
	for _, row := range rows {
		split := strings.Split(row, " ")
		var numbers []int
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			numbers = append(numbers, n)
		}
		totalSum += getHistoryPrediction(numbers)
	}
	fmt.Println(totalSum)
}
