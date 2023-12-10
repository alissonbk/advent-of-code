package partone

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

const (
	NORTH = iota
	WEST
	SOUTH
	EAST
)

type Pos struct {
	x                    int
	y                    int
	directionFromLastPos int // NWSE
}

func (p Pos) getCurrentChar(rows *[][]string) string {
	return (*rows)[p.y][p.x]
}

func (p Pos) isEqual(pos Pos) bool {
	return p.x == pos.x && p.y == pos.y
}

func findS(rows []string) Pos {
	x := 0
	y := 0
	for idx, r := range rows {
		if strings.Contains(r, "S") {
			y = idx
			split := strings.Split(r, "")
			for i, c := range split {
				if c == "S" {
					x = i
				}
			}
		}
	}
	return Pos{x: x, y: y}
}

func splitRows(rows []string) [][]string {
	var split [][]string
	for _, row := range rows {
		split = append(split, strings.Split(row, ""))
	}
	return split
}

// fixme
func getAvaliablePositions(pos Pos, xLength int, yLength int, currentChar string) []Pos {
	var avaliablePos []Pos
	east := Pos{pos.x + 1, pos.y, EAST}
	south := Pos{pos.x, pos.y + 1, SOUTH}
	west := Pos{pos.x - 1, pos.y, WEST}
	north := Pos{pos.x, pos.y - 1, NORTH}

	if pos.y == 0 {
		if pos.x == 0 {
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, south)
		} else if pos.x == xLength-1 {
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, south)
		} else {
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, south)
		}
	} else if pos.y == yLength-1 {
		if pos.x == 0 {
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, north)
		} else if pos.x == xLength-1 {
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, north)
		} else {
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, north)
		}
	} else {
		if pos.x == 0 {
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, north)
			avaliablePos = append(avaliablePos, south)
		} else if pos.x == xLength-1 {
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, south)
			avaliablePos = append(avaliablePos, north)
		} else {
			avaliablePos = append(avaliablePos, west)
			avaliablePos = append(avaliablePos, east)
			avaliablePos = append(avaliablePos, south)
			avaliablePos = append(avaliablePos, north)
		}
	}
	if currentChar == "S" {
		return avaliablePos
	} else if currentChar == "." {
		return nil
	} else {
		var filtered []Pos
		for _, ap := range avaliablePos {
			if currentChar == "|" && (ap.directionFromLastPos == NORTH || ap.directionFromLastPos == SOUTH) {
				filtered = append(filtered, ap)
				continue
			}
			if currentChar == "-" && (ap.directionFromLastPos == EAST || ap.directionFromLastPos == WEST) {
				filtered = append(filtered, ap)
				continue
			}
			if currentChar == "L" && (ap.directionFromLastPos == NORTH || ap.directionFromLastPos == EAST) {
				filtered = append(filtered, ap)
				continue
			}
			if currentChar == "J" && (ap.directionFromLastPos == NORTH || ap.directionFromLastPos == WEST) {
				filtered = append(filtered, ap)
				continue
			}
			if currentChar == "7" && (ap.directionFromLastPos == SOUTH || ap.directionFromLastPos == WEST) {
				filtered = append(filtered, ap)
				continue
			}
			if currentChar == "F" && (ap.directionFromLastPos == SOUTH || ap.directionFromLastPos == EAST) {
				filtered = append(filtered, ap)
				continue
			}
		}
		return filtered
	}
}

func getStartingPoints(avaliablePositions []Pos, rows [][]string, xLen int, yLen int) []Pos {
	var filtered []Pos
	for _, ap := range avaliablePositions {
		apChar := ap.getCurrentChar(&rows)
		newPositions := getAvaliablePositions(ap, xLen, yLen, apChar)
		for _, np := range newPositions {
			if np.getCurrentChar(&rows) == "S" { // is pointing to S
				filtered = append(filtered, ap)
				break
			}
		}
	}
	return filtered
}

func removeSFromPositions(pos []Pos, rowSplit *[][]string) []Pos {
	var newP []Pos
	for _, p := range pos {
		if p.getCurrentChar(rowSplit) != "S" {
			newP = append(newP, p)
		}
	}
	return newP
}

func Run() {
	rows := utils.GetSliceFromFile("/day10/partone/input.txt")
	sPos := findS(rows)
	yLen := len(rows)
	xLen := len(rows[0])
	rowSplit := splitRows(rows)
	reachedSAgain := false
	startingPoints := getStartingPoints(getAvaliablePositions(sPos, xLen, yLen, "S"), rowSplit, xLen, yLen)
	startPoint := startingPoints[0]
	endPoint := startingPoints[1]
	currentPos := startPoint
	lastPos := currentPos
	firstIteration := true
	totalIterations := 0
	for reachedSAgain == false {
		currentChar := currentPos.getCurrentChar(&rowSplit)
		aps := getAvaliablePositions(currentPos, xLen, yLen, currentChar)

		if firstIteration {
			aps = removeSFromPositions(aps, &rowSplit)
			currentPos = aps[0]
			firstIteration = false
			totalIterations += 1
			continue
		}
		if currentPos.isEqual(endPoint) {
			reachedSAgain = true
			break
		}
		for _, ap := range aps {
			if !ap.isEqual(lastPos) {
				lastPos = currentPos
				currentPos = ap
				totalIterations += 1
				break
			}
		}
	}
	// reaching S from starting point + end point
	totalIterations += 2
	fmt.Println("total Iterations: ", totalIterations/2)
}
