package partone

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Pos struct {
	y int
	x int
}

func (p Pos) isEqual(pos Pos) bool {
	return p.x == pos.x && p.y == pos.y
}

func incrementColumnsToAppend(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] + i
	}
}

func Run() {
	rows := utils.GetSliceFromFile("/day11/partone/input.txt")
	var galaxiesPositions []Pos
	var taggedToAppendCol []int
	rowsLen := len(rows)
	totalSum := 0
	for y := 0; y < rowsLen; y++ {
		foundGalaxyInRow := false
		for x, c := range rows[y] {
			if string(c) == "#" {
				foundGalaxyInRow = true
			}
			if y == 0 {
				foundGalaxyInCol := false
				for j := 0; j < len(rows); j++ {
					if string(rows[j][x]) == "#" {
						foundGalaxyInCol = true
					}
				}
				if !foundGalaxyInCol {
					taggedToAppendCol = append(taggedToAppendCol, x)
				}
			}

		}
		if !foundGalaxyInRow {
			utils.InsertAtIndex(&rows, y, rows[y])
			rowsLen++
			y++
		}
	}

	incrementColumnsToAppend(&taggedToAppendCol)
	for y := 0; y < rowsLen; y++ {
		for x := 0; x < len(rows[0]); x++ {
			for _, tp := range taggedToAppendCol {
				if x == tp {
					rowSplit := strings.Split(rows[y], "")
					utils.InsertAtIndex(&rowSplit, x, ".")
					rows[y] = strings.Join(rowSplit, "")
				}
			}
		}
	}

	for y, row := range rows {
		for x, c := range row {
			if string(c) == "#" {
				galaxiesPositions = append(galaxiesPositions, Pos{y, x})
			}
		}
	}

	gp := galaxiesPositions[0]

	for len(galaxiesPositions) > 0 {
		for _, g := range galaxiesPositions {
			if !gp.isEqual(g) {
				sum := 0
				if gp.x >= g.x {
					sum += gp.x - g.x
				} else {
					sum += g.x - gp.x
				}
				if gp.y >= g.y {
					sum += gp.y - g.y
				} else {
					sum += g.y - gp.y
				}
				totalSum += sum
			}
		}

		galaxiesPositions = utils.RemoveIndex(galaxiesPositions, 0)
		if len(galaxiesPositions) > 0 {
			gp = galaxiesPositions[0]
		}
	}

	fmt.Println(totalSum)
}
