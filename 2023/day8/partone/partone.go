package partone

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type MapNode struct {
	location string
	leftDst  string
	rightDst string
}

type SearchedLocation struct {
	location string
	index    int
}

func getMapNode(str string) MapNode {
	var mapNode MapNode

	split := strings.Split(str, " ")
	mapNode.location = split[0]
	mapNode.leftDst = strings.ReplaceAll(strings.ReplaceAll(split[2], "(", ""), ",", "")
	mapNode.rightDst = strings.ReplaceAll(split[3], ")", "")
	return mapNode
}

func jmpCurrentInstruction(instructions []string, currentInstructionPos *int) {
	if *currentInstructionPos < len(instructions)-1 {
		*currentInstructionPos += 1
	} else {
		*currentInstructionPos = 0
	}
}

func findIndexToStart(rows []string) int {
	for idx, r := range rows {
		if strings.HasPrefix(r, "AAA") {
			return idx
		}
	}
	return 2
}

func Run() {
	rows := utils.GetSliceFromFile("/day8/partone/input.txt")
	instructions := strings.Split(rows[0], "")
	totalSteps := 0
	foundZ := false
	index := findIndexToStart(rows)
	currentInstructionPos := 0
	locationToSearch := ""
	var alreadySearchedLocations []SearchedLocation
	for foundZ == false {
		node := getMapNode(rows[index])
		ignoreIndex := false
		alreadySearchedLocations = append(alreadySearchedLocations, SearchedLocation{node.location, index})
		if node.location == locationToSearch || locationToSearch == "" {
			if node.location == "ZZZ" {
				foundZ = true
			} else {
				if instructions[currentInstructionPos] == "L" {
					locationToSearch = node.leftDst
				}
				if instructions[currentInstructionPos] == "R" {
					locationToSearch = node.rightDst
				}
				for _, l := range alreadySearchedLocations {
					if l.location == locationToSearch {
						index = l.index
						ignoreIndex = true
						alreadySearchedLocations = nil
					}
				}
				totalSteps += 1
			}
			jmpCurrentInstruction(instructions, &currentInstructionPos)
		}
		if !ignoreIndex && index < len(rows)-1 {
			index += 1
		} else if index == len(rows)-1 {
			index = 2
		}
	}

	fmt.Println("Total Steps: ", totalSteps)

}
