package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strconv"
)

type ValidPos struct {
	x int
	y int
}

type ValidPositions []ValidPos

type NumberInfo struct {
	y          int
	startIdx   int
	endIdx     int
	realNumber int
}

type NumberAlreadyTaken struct {
	rowIdx   int
	firstPos int
	lastPos  int
}

type AllNumbersAlreadyTaken []NumberAlreadyTaken

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func isGear(b byte) bool {
	return b == 42
}

type IndexesValidPosition struct {
	idx   int
	vpIdx int
}

func (vps *ValidPositions) removeDuplicateds() {
	var rows []int
	var rowsWithDuplicateds []int
	var idxA []IndexesValidPosition
	var idxB []IndexesValidPosition
	for _, vp := range *vps {
		rows = append(rows, vp.y)
		counter := 0
		for _, n := range rows {
			if vp.y == n {
				counter += 1
			}
		}
		if counter > 1 {
			alreadyExists := false
			for _, rw := range rowsWithDuplicateds {
				if rw == vp.y {
					alreadyExists = true
				}
			}
			if !alreadyExists {
				rowsWithDuplicateds = append(rowsWithDuplicateds, vp.y)
			}

		}

	}
	if len(rowsWithDuplicateds) > 2 {
		panic("Should not have more than 2 rows with duplicateddnumbers!")
	}
	for i, vp := range *vps {
		if len(rowsWithDuplicateds) > 0 && vp.y == rowsWithDuplicateds[0] {
			ivp := IndexesValidPosition{idx: vp.x, vpIdx: i}
			idxA = append(idxA, ivp)
		}
		if len(rowsWithDuplicateds) > 1 && vp.y == rowsWithDuplicateds[1] {
			ivp := IndexesValidPosition{idx: vp.x, vpIdx: i}
			idxB = append(idxB, ivp)
		}
	}

	if len(idxA) > 3 || len(idxB) > 3 {
		panic("Should not have more than 3 indexes for a number duplication check")
	}

	removedCount := 0
	if len(idxA) >= 2 {
		if math.Abs(float64(idxA[1].idx-idxA[0].idx)) == 1 {
			*vps = utils.RemoveIndex(*vps, idxA[1].vpIdx)
			removedCount += 1
		}

		if len(idxA) > 2 && (math.Abs(float64(idxA[1].idx-idxA[2].idx)) == 1 || math.Abs(float64(idxA[1].idx-idxA[0].idx)) == 1) {
			*vps = utils.RemoveIndex(*vps, idxA[1].vpIdx)
			removedCount += 1
		}
	}

	if len(idxB) >= 2 {
		if math.Abs(float64(idxB[1].idx-idxB[0].idx)) == 1 {
			*vps = utils.RemoveIndex(*vps, idxB[0].vpIdx-removedCount)
			removedCount += 1
		}

		if len(idxB) > 2 && (math.Abs(float64(idxB[1].idx-idxB[2].idx)) == 1 || math.Abs(float64(idxB[1].idx-idxB[0].idx)) == 1) {
			*vps = utils.RemoveIndex(*vps, idxB[1].vpIdx-removedCount)
		}
	}

}

func (anat *AllNumbersAlreadyTaken) isNumberTaken(validPosition ValidPos) bool {
	for _, nt := range *anat {
		if nt.rowIdx == validPosition.y {
			for i := nt.firstPos; i < nt.lastPos; i++ {
				if validPosition.x == i {
					return true
				}
			}
		}
	}
	return false
}

func sumValidGear(
	validPositions [2]ValidPos, rows *[][]byte,
	totalSum *int) {
	var allNumbersPositions [][]int
	var finalNumbers [2]int

	for i, vp := range validPositions {
		var idxsWithNumbers []int
		for idx, c := range (*rows)[vp.y] {
			if isNumber(c) {
				idxsWithNumbers = append(idxsWithNumbers, idx)
			} else {
				idxsWithNumbers = append(idxsWithNumbers, -1)
			}
		}
		var numberPositions []int
		for idx, idxWithNumber := range idxsWithNumbers {
			if idxWithNumber != -1 {
				numberPositions = append(numberPositions, idxWithNumber)
			}
			if (idxWithNumber == -1 || idx == len(idxsWithNumbers)-1) && len(numberPositions) > 0 {
				for _, np := range numberPositions {
					if np == vp.x {
						allNumbersPositions = append(allNumbersPositions, numberPositions)
					}
				}
				numberPositions = nil
			}
		}
		var numString string
		for _, n := range allNumbersPositions[i] {
			numString += string((*rows)[vp.y][n])
		}
		conv, _ := strconv.Atoi(numString)
		finalNumbers[i] = conv

	}

	*totalSum += finalNumbers[0] * finalNumbers[1]
}

func Run() {
	// Solution would be faster by doing stuff while reading the bytes from the file
	//"not appending to slice and reading it again"
	rows := utils.GetByteArrSliceFromFile("/day3/parttwo/input.txt")

	totalSum := 0
	//var allTakenNumbers AllNumbersAlreadyTaken
	colLen := len(rows)
	for idx, row := range rows {
		rowLen := len(row)
		for i, b := range row {
			if isGear(b) {
				// check all postions arround
				if idx == 0 { //first row
					if i == 0 { //first el
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else if i == rowLen-1 { //last el
						var numbersAroundGears ValidPositions

						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else { //mid el
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					}

				} else if idx == colLen-1 { // last row
					if i == 0 { // first el
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{x: i, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else if i == rowLen-1 { //last el
						var numbersAroundGears ValidPositions

						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{x: i, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else { // mid el
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{x: i, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					}

				} else { // from row 2 to len-1
					if i == 0 { //first el
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{i, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i+1]) {
							vp := ValidPos{i + 1, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else if i == rowLen-1 { //last el
						var numbersAroundGears ValidPositions

						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{x: i, y: idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i-1]) {
							vp := ValidPos{i - 1, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					} else { // check every thing - mid element -
						var numbersAroundGears ValidPositions

						if isNumber(row[i+1]) {
							vp := ValidPos{x: i + 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(row[i-1]) {
							vp := ValidPos{x: i - 1, y: idx}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i]) {
							vp := ValidPos{i, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i-1]) {
							vp := ValidPos{i - 1, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx-1][i+1]) {
							vp := ValidPos{i + 1, idx - 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i]) {
							vp := ValidPos{x: i, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i-1]) {
							vp := ValidPos{x: i - 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}
						if isNumber(rows[idx+1][i+1]) {
							vp := ValidPos{x: i + 1, y: idx + 1}
							numbersAroundGears = append(numbersAroundGears, vp)
						}

						numbersAroundGears.removeDuplicateds()
						if len(numbersAroundGears) == 2 {
							validPositions := [2]ValidPos{numbersAroundGears[0], numbersAroundGears[1]}
							sumValidGear(validPositions, &rows, &totalSum)
						}
					}
				}
			}
		}

	}
	fmt.Println("\n totalSum: ", totalSum)
}
