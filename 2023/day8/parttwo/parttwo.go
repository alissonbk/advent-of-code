package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strings"
	"sync"
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

type StepsByGoroutines struct {
	id              int
	stepsTookToFind []int
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

func findIndexesToStart(rows []string) []int {
	var indexes []int
	for idx, r := range rows {
		if idx > 1 {
			split := strings.Split(r, " ")
			sp := strings.Split(split[0], "")
			if sp[2] == "A" {
				indexes = append(indexes, idx)
			}
		}

	}
	return indexes
}

func isEndingWithZ(str string) bool {
	split := strings.Split(str, "")
	if split[2] == "Z" {
		return true
	}
	return false
}

func appendInOrder(c chan int, totalStepsReference []int, maxLength int, wg *sync.WaitGroup) {
	var totalSteps []int
	for {
		if len(totalSteps) == maxLength {
			for i := 0; i < len(totalStepsReference); i++ {
				totalStepsReference[i] = totalSteps[i]
			}
			close(c)
			wg.Done()
		}
		step := <-c
		if totalSteps == nil {
			totalSteps = append(totalSteps, step)
		} else {
			for idx, st := range totalSteps {
				if step < st {
					utils.InsertAtIndex(&totalSteps, idx, step)

				} else if idx == len(totalSteps)-1 {
					totalSteps = append(totalSteps, step)
				}
			}
		}

	}

}

func search(instructions []string, rows []string, index int, stepChan chan<- int) {
	foundZ := false
	currentInstructionPos := 0
	locationToSearch := ""
	steps := 0
	var alreadySearchedLocations []SearchedLocation
	for foundZ == false {
		node := getMapNode(rows[index])
		ignoreIndex := false
		alreadySearchedLocations = append(alreadySearchedLocations, SearchedLocation{node.location, index})
		if node.location == locationToSearch || locationToSearch == "" {
			if isEndingWithZ(node.location) {
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
				steps += 1
			}
			jmpCurrentInstruction(instructions, &currentInstructionPos)
		}
		if !ignoreIndex && index < len(rows)-1 {
			index += 1
		} else if index == len(rows)-1 {
			index = 2
		}
	}
	fmt.Println("Goroutine found in: ", steps, "steps")
	stepChan <- steps
}

func Run() {
	rows := utils.GetSliceFromFile("/day8/parttwo/input.txt")
	instructions := strings.Split(rows[0], "")
	indexes := findIndexesToStart(rows)
	totalSteps := make([]int, len(indexes))
	stepChan := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go appendInOrder(stepChan, totalSteps, len(indexes), &wg)
	for _, idx := range indexes {
		go search(instructions, rows, idx, stepChan)
	}
	wg.Wait()

	a := totalSteps[0]
	utils.RemoveIndex(totalSteps, 0)
	b := totalSteps[0]
	utils.RemoveIndex(totalSteps, 0)
	finalSum := utils.LCM(a, b, totalSteps...)

	fmt.Println(finalSum)
}
