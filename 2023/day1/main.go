package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

// ------------------PART ONE-----------------------
func partOne() {
	rows := utils.GetSliceFromFile("/day1/input.txt")
	var totalSum int

	for _, row := range rows {
		var rowNumbers []string
		for _, letter := range strings.Split(row, "") {
			number, err := strconv.Atoi(letter)
			if err == nil {
				rowNumbers = append(rowNumbers, strconv.Itoa(number))
			}
		}
		var numberAsString = rowNumbers[0] + rowNumbers[len(rowNumbers)-1]
		n, _ := strconv.Atoi(numberAsString)
		totalSum += n
	}

	fmt.Println("part one result: ", totalSum)
}

// ------------------PART TWO-----------------------
type foundNumbers struct {
	n     string
	index int
}

func getSpelledNumbers() []string {
	return []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine"}
}

func getRowSum(first foundNumbers, last foundNumbers) int {
	n, _ := strconv.Atoi(first.n + last.n)
	return n
}

// Using recursion solves some exclusive edge case like -> a5awimzeighteightwol
func addSpelledNumbersRecursively(fn *[]foundNumbers, row string, sn string, idx int, lastWordIndexTakenLst *[]int) {
	if strings.Contains(row, sn) {
		firstIndex := strings.Index(row, sn)
		*lastWordIndexTakenLst = append(*lastWordIndexTakenLst, firstIndex+(len(sn)-1))
		*fn = append(*fn, foundNumbers{
			strconv.Itoa(idx + 1),
			firstIndex})

		newRow := strings.Replace(row, sn, "", 1)
		if strings.Contains(newRow, sn) {
			var filledRow string
			for i := 0; i < len(sn); i++ {
				filledRow += " "
			}
			filledRow += newRow
			addSpelledNumbersRecursively(fn, filledRow, sn, idx, lastWordIndexTakenLst)
		}
	}
}

func partTwo() {
	rows := utils.GetSliceFromFile("day1/input.txt")
	var totalSum int
	var lastWordIndexTakenLst []int
	spelledNumbers := getSpelledNumbers()
	for _, row := range rows {
		var fn []foundNumbers
		// Search for spelled numbers
		for i, sn := range spelledNumbers {
			addSpelledNumbersRecursively(&fn, row, sn, i, &lastWordIndexTakenLst)
		}
		// Search for 'real' integers
		for i, letter := range strings.Split(row, "") {
			_, err := strconv.Atoi(letter)
			if err == nil {
				fn = append(fn, foundNumbers{letter, i})
			}
		}
		lastWordIndexTakenLst = nil

		var lowerRowIdx int
		var highestRowIdx int
		// It's not sorted because the loop through `getSpelledNumbers()`
		for i, n := range fn {
			if n.index <= fn[lowerRowIdx].index {
				lowerRowIdx = i
			}
			if n.index > fn[highestRowIdx].index {
				highestRowIdx = i
			}
		}

		totalSum += getRowSum(fn[lowerRowIdx], fn[highestRowIdx])
	}

	fmt.Println("part two result: ", totalSum)
}

func main() {
	partOne()
	partTwo()
}
