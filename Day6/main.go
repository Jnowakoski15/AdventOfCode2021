package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type lanternFish struct {
	timer int
}

func (l *lanternFish) isZero() bool {
	return l.timer == 0
}

func (l *lanternFish) decreaseTimer() {
	l.timer--
}

func (l *lanternFish) resetTimer() {
	l.timer = 6
}

func main() {
	part1 := lanterFishSimulation("input.txt", 80)
	part2 := lanterFishSimulation2("input.txt", 256)
	fmt.Printf("Fish after 80 days:  %v \n", part1)
	fmt.Printf("Fish after 256 days : %v \n", part2)

}

func lanterFishSimulation2(s string, days int) int {
	fishes := readInFileToIntArray(s)
	fishMap := make(map[int]int)
	fishTimerMax := 8
	//init fish Map
	for i := 0; i <= fishTimerMax; i++ {
		fishMap[i] = 0
	}

	for _, fish := range fishes {
		fishMap[fish.timer]++
	}

	for i := 0; i < days; i++ {
		spawnCount := 0
		for timer := 0; timer <= fishTimerMax; timer++ {
			if timer == 0 {
				spawnCount = fishMap[timer]
			}
			fishMap[timer] = fishMap[timer+1]
		}
		fishMap[6] += spawnCount
		fishMap[fishTimerMax] = spawnCount
	}

	numOfFish := 0

	for _, val := range fishMap {
		numOfFish += val
	}

	return numOfFish
}

func lanterFishSimulation(s string, days int) int {
	fishes := readInFileToIntArray(s)

	for i := 0; i < days; i++ {
		spawnCount := 0
		for _, fish := range fishes {
			if fish.isZero() {
				spawnCount++
				fish.resetTimer()
			} else {
				fish.decreaseTimer()
			}

		}
		// spawnFish
		fishes = spawnFish(spawnCount, fishes)
	}

	return len(fishes)
}

func spawnFish(spawnCount int, fishes []*lanternFish) []*lanternFish {
	defaultFishTimer := 8
	for i := 0; i < spawnCount; i++ {
		fishes = append(fishes, &lanternFish{defaultFishTimer})
	}
	return fishes
}

func readInFileToIntArray(name string) []*lanternFish {
	output := []*lanternFish{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	numArry := strings.Split(lines[0], ",")

	for _, numStr := range numArry {
		val, _ := strconv.Atoi(numStr)
		fish := lanternFish{val}
		output = append(output, &fish)
	}
	return output
}
