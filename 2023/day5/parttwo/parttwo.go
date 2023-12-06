package parttwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
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

func produceSeeds(oldSeeds []int, seedsQueue chan<- *uint32, wg *sync.WaitGroup) {
	defer close(seedsQueue)
	for idx := 0; idx < len(oldSeeds); idx++ {
		if idx == len(oldSeeds)-1 {
			break
		}

		if idx%2 == 0 {
			iterator := oldSeeds[idx+1]
			for i := 0; i < iterator; i++ {
				n := uint32(oldSeeds[idx] + i)
				seedsQueue <- &n
			}
		}
	}
	wg.Done()
}

func consumeSeeds(seedsQueue <-chan *uint32, almanacParts AlmanacParts, minFinalLocation *int) {
	var destination int
	for seed := range seedsQueue {
		destination = int(*seed)
		for _, ap := range almanacParts {
			for _, p := range *ap.parts {
				sourceSum := p.srcRangeStart + (p.length - 1)
				if destination >= p.srcRangeStart && destination <= sourceSum { // found mapping
					destination = p.dstRangeStart - p.srcRangeStart + destination
					break
				}
			}
		}
		if destination < *minFinalLocation || *minFinalLocation == -1 {
			*minFinalLocation = destination
		}
	}

}

func Run() {
	rows := utils.GetSliceFromFile("/day5/parttwo/input.txt")
	seedsStr, _ := strings.CutPrefix(rows[0], "seeds: ")
	oldSeeds := getSliceFromStringRow(seedsStr)
	seedsQueue := make(chan *uint32, 10) // 16bits would break because of number size
	var almanacParts AlmanacParts
	minFinalLocation := -1
	wg := sync.WaitGroup{}
	wg.Add(1)
	now := time.Now()

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

	go produceSeeds(oldSeeds, seedsQueue, &wg)
	go consumeSeeds(seedsQueue, almanacParts, &minFinalLocation)
	wg.Wait()

	fmt.Println("minimun location ", minFinalLocation)
	fmt.Println("took ", time.Now().Sub(now).Seconds(), " seconds")
}
