package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1 := depthMesaureIncrease("input.txt")
	part2 := slideSumIncrease("input.txt")
	fmt.Printf("Depth changes: %v \n", part1)
	fmt.Printf("Sliding sum changes: %v \n", part2)
}

func depthMesaureIncrease(input string) int {
	lines := readInFileToIntArray(input)
	i := 0
	output := 0
	for i < len(lines)-1 {
		if lines[i] < lines[i+1] {
			output++
		}
		i++
	}

	return output
}

func slideSumIncrease(input string) int {
	lines := readInFileToIntArray(input)
	i := 0
	output := 0
	mapOfValues := make(map[int]int)

	for i < len(lines) {
		j := i
		counter := 0

		for counter < 3 && j < len(lines) {
			mapOfValues[i] += lines[j]
			j++
			counter++
		}
		i++
	}
	i = 0
	for i < len(mapOfValues)-1 {
		if mapOfValues[i] < mapOfValues[i+1] {
			output++
		}
		i++
	}

	return output
}

func readInFileToIntArray(name string) []int {
	output := []int{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		output = append(output, val)
	}
	return output
}
