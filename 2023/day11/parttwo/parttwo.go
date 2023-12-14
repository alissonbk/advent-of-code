package parttwo

import (
	"aoc2023/utils"
	"fmt"
)

type Pos struct {
	y int
	x int
}

func (p Pos) isEqual(pos Pos) bool {
	return p.x == pos.x && p.y == pos.y
}

func howManyPointsToIncrease(startPos int, destPos int, toAppend []int, destItsBigger bool) int {
	sum := 0
	for _, ta := range toAppend {
		if destItsBigger && ta >= startPos && ta <= destPos {
			sum += 1
		}
		if !destItsBigger && ta >= destPos && ta <= startPos {
			sum += 1
		}
	}
	return sum * 999999
}

func Run() {
	rows := utils.GetSliceFromFile("/day11/parttwo/input.txt")
	var galaxiesPositions []Pos
	rowsLen := len(rows)
	totalSum := 0
	var yToAppend []int
	var xToAppend []int
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
					xToAppend = append(xToAppend, x)
				}
			}

		}
		if !foundGalaxyInRow {
			yToAppend = append(yToAppend, y)
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
					sum += (gp.x - g.x) + howManyPointsToIncrease(gp.x, g.x, xToAppend, false)
				} else {
					sum += (g.x - gp.x) + howManyPointsToIncrease(gp.x, g.x, xToAppend, true)
				}
				if gp.y >= g.y {
					sum += (gp.y - g.y) + howManyPointsToIncrease(gp.y, g.y, yToAppend, false)
				} else {
					sum += (g.y - gp.y) + howManyPointsToIncrease(gp.y, g.y, yToAppend, true)
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
