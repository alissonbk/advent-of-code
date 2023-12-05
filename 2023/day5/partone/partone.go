package partone

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Part struct {
	dstRangeStart int
	srcRangeStart int
	length        int
}

type AlmanacPart struct {
	name  string
	parts *[]Part
}

type AlmanacParts []AlmanacPart

func (aps *AlmanacParts) appendPart(name string, part Part) {
	found := false
	for _, ap := range *aps {
		if ap.name == name {
			found = true
			*ap.parts = append(*ap.parts, part)
		}
	}
	if !found {
		newAlmanacPart := AlmanacPart{name: name, parts: &[]Part{part}}
		*aps = append(*aps, newAlmanacPart)
	}
}

func getSliceFromStringRow(str string) []int {
	var slice []int
	split := strings.Split(str, " ")
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		slice = append(slice, n)
	}
	return slice
}

func Run() {
	rows := utils.GetSliceFromFile("/day5/partone/input.txt")
	seedsStr, _ := strings.CutPrefix(rows[0], "seeds: ")
	seeds := getSliceFromStringRow(seedsStr)
	var almanacParts AlmanacParts
	var destination int
	minFinalLocation := -1

	for idx := 1; idx < len(rows); idx++ {
		if strings.Contains(rows[idx], "map:") {
			name := strings.Replace(rows[idx], " map:", "", -1)
			idx += 1
			for rows[idx] != "" {
				numbers := getSliceFromStringRow(rows[idx])
				part := Part{
					dstRangeStart: numbers[0],
					srcRangeStart: numbers[1],
					length:        numbers[2],
				}
				almanacParts.appendPart(name, part)
				if idx >= len(rows)-1 {
					break
				}
				idx += 1
			}
		}
	}

	for _, seed := range seeds {
		destination = seed
		for _, ap := range almanacParts {
			for _, p := range *ap.parts {
				sourceSum := p.srcRangeStart + (p.length - 1)
				if destination >= p.srcRangeStart && destination <= sourceSum { // found mapping
					destination = p.dstRangeStart - p.srcRangeStart + destination
					break
				}
			}
		}
		if destination < minFinalLocation || minFinalLocation == -1 {
			minFinalLocation = destination
		}
	}

	fmt.Println("minimun location ", minFinalLocation)

}
