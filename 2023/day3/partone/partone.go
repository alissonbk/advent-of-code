package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
)

type ValidPos struct {
	x int
	y int
}

type NumberInfo struct {
	y           int
	startIdx    int
	endIdx      int
	realNumber  int
	isValidated bool
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func isDot(b byte) bool {
	return b == 46
}

func isSpecialCharacter(b byte) bool {
	return !isNumber(b) && !isDot(b)
}

func setValidNumber(numbers *[]NumberInfo, x int, y int, row []byte) int {
	for idx, n := range *numbers {
		if n.y == y && n.isValidated == false {
			for i := n.startIdx; i <= n.endIdx; i++ {
				if i == x {
					(*numbers)[idx].isValidated = true
					return sumRealNumber(&n, row)
				}
			}
		}
	}
	return 0
}

func sumRealNumber(ni *NumberInfo, row []byte) int {
	var bytearr []byte
	var str string

	for i := ni.startIdx; i <= ni.endIdx; i++ {
		bytearr = append(bytearr, row[i])
	}
	for _, b := range bytearr {
		str += string(b)
	}
	n, _ := strconv.Atoi(str)
	return n
}

func checkNumberWasValidatedBefore(ni *NumberInfo, posIsValid *[]ValidPos) bool {
	for _, vp := range *posIsValid {
		if ni.y == vp.y {
			for i := ni.startIdx; i <= ni.endIdx; i++ {
				if i == vp.x {
					return true
				}
			}
		}
	}
	return false
}

func Run() {
	// Solution would be faster by doing stuff while reading the bytes from the file
	//"not appending to slice and reading it again"
	rows := utils.GetByteArrSliceFromFile("/day3/partone/input.txt")

	totalSum := 0
	var numbers []NumberInfo
	var posIsValid []ValidPos
	colLen := len(rows)
	for idx, row := range rows {
		lastWasNumber := false
		firstNumberIdx := -1
		rowLen := len(row)
		for i, b := range row {
			if isNumber(b) {
				if !lastWasNumber {
					firstNumberIdx = i
				}
				if i == rowLen-1 {
					ni := NumberInfo{
						y:        idx,
						startIdx: firstNumberIdx,
						endIdx:   i}
					isValid := checkNumberWasValidatedBefore(&ni, &posIsValid)
					if isValid {
						ni.isValidated = true
						totalSum += sumRealNumber(&ni, row)
					}

					numbers = append(numbers, ni)
					firstNumberIdx = -1
				}
				lastWasNumber = true

			} else {
				if lastWasNumber {
					ni := NumberInfo{
						y:        idx,
						startIdx: firstNumberIdx,
						endIdx:   i - 1}
					isValid := isSpecialCharacter(b) || checkNumberWasValidatedBefore(&ni, &posIsValid)
					if isValid {
						ni.isValidated = true
						totalSum += sumRealNumber(&ni, row)
					}

					numbers = append(numbers, ni)
					firstNumberIdx = -1
				}
				lastWasNumber = false
			}
			if isSpecialCharacter(b) {
				// check all postions arround
				if idx == 0 { //first row
					if i == 0 { //first el
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx + 1})
						}
					} else if i == rowLen-1 { //last el
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i-1]) {
							posIsValid = append(posIsValid, ValidPos{x: i - 1, y: idx + 1})
						}
					} else { //mid el
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}
						if isNumber(rows[idx+1][i-1]) {
							posIsValid = append(posIsValid, ValidPos{x: i - 1, y: idx + 1})
						}
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx + 1})
						}
					}

				} else if idx == colLen-1 { // last row
					if i == 0 { // first el
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}
						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i+1]) {
							totalSum += setValidNumber(&numbers, i+1, idx-1, rows[idx-1])
						}
					} else if i == rowLen-1 { //last el
						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx+1, rows[idx-1])
						}
						if isNumber(rows[idx+1][i-1]) {
							posIsValid = append(posIsValid, ValidPos{x: i - 1, y: idx + 1})
						}
					} else { // mid el
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}
						if isNumber(rows[idx-1][i-1]) {
							totalSum += setValidNumber(&numbers, i-1, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i+1]) {
							totalSum += setValidNumber(&numbers, i+1, idx-1, rows[idx-1])
						}
					}
				} else { // from row 2 to len-1
					if i == 0 { //first el
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx + 1})
						}
						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i+1]) {
							totalSum += setValidNumber(&numbers, i+1, idx-1, rows[idx-1])
						}
					} else if i == rowLen-1 { //last el
						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i-1]) {
							totalSum += setValidNumber(&numbers, i-1, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i-1]) {
							posIsValid = append(posIsValid, ValidPos{x: i - 1, y: idx + 1})
						}
					} else { // check every thing - mid element -
						if isNumber(row[i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx})
						}

						if isNumber(rows[idx-1][i]) {
							totalSum += setValidNumber(&numbers, i, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i-1]) {
							totalSum += setValidNumber(&numbers, i-1, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx-1][i+1]) {
							totalSum += setValidNumber(&numbers, i+1, idx-1, rows[idx-1])
						}
						if isNumber(rows[idx+1][i]) {
							posIsValid = append(posIsValid, ValidPos{x: i, y: idx + 1})
						}
						if isNumber(rows[idx+1][i-1]) {
							posIsValid = append(posIsValid, ValidPos{x: i - 1, y: idx + 1})
						}
						if isNumber(rows[idx+1][i+1]) {
							posIsValid = append(posIsValid, ValidPos{x: i + 1, y: idx + 1})
						}
					}
				}
			}
		}

	}
	fmt.Println("\n totalSum: ", totalSum)
}

/*
.#959....732...724*328...*..59...
............*..........35........
5...........922.....+....550.....
..................902.....*......
*/
