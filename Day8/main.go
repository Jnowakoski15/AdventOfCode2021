package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type numberSet struct {
	left  []string
	right []string
}

func main() {
	part1 := part1("input.txt")
	//part2 := part2("input.txt")
	fmt.Printf("Number of 1,4,7,8:  %v \n", part1)
	//fmt.Printf("Least Fuel Used:  %v \n", part1)
}

func part2(s string) int {
	outputCount := 0
	defaultLetterMap := make(map[string]string)
	defaultLetterMap["abcefg"] = "0" //6
	defaultLetterMap["cf"] = "1"
	defaultLetterMap["acdeg"] = "2"   //5
	defaultLetterMap["acdfg"] = "3"   //5
	defaultLetterMap["bdcf"] = "4"    //4
	defaultLetterMap["abdfg"] = "5"   //5
	defaultLetterMap["abdefg"] = "6"  //6
	defaultLetterMap["acf"] = "7"     //3
	defaultLetterMap["abcdefg"] = "8" // 7
	defaultLetterMap["abcdfg"] = "9"  // 6
	numberMap := make(map[int]int)
	// 3 letters = 7
	// 2 letters = 1
	// 4 letters = 4
	// 7 letters = 8
	numberMap[2] = 1
	numberMap[4] = 4
	numberMap[3] = 7
	numberMap[7] = 8

	letterMap := make(map[string]string)
	numberSets := readInFileToIntArray(s)
	for _, line := range numberSets {
		//First Pass setup
		for _, rightVal := range line.right {
			letterCount := len(rightVal)
			switch letterCount {
			case 3:
				letterMap[rightVal] = "7"
			case 2:
				letterMap[rightVal] = "1"
			case 4:
				letterMap[rightVal] = "4"
			case 7:
				letterMap[rightVal] = "8"
			}
		}

		//Second pass
		// 9 contains a 7
		// 3 contains a 1
		// 2 doesn't contain a 1
		// 

	}
	return outputCount
}

func part1(s string) int {
	outputCount := 0
	// 3 letters = 7
	// 2 letters = 1
	// 4 letters = 4
	// 7 letters = 8
	numberMap := make(map[int]int)
	numberMap[2] = 1
	numberMap[4] = 4
	numberMap[3] = 7
	numberMap[7] = 8

	numberSets := readInFileToIntArray(s)
	for _, line := range numberSets {
		for _, rightVal := range line.right {
			letterCount := len(rightVal)
			_, isFound := numberMap[letterCount]
			if isFound {
				outputCount++
			}
		}
	}
	return outputCount
}

func readInFileToIntArray(name string) []*numberSet {
	output := []*numberSet{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")

		left := strings.Split(splitLine[0], " ")
		right := strings.Split(splitLine[1], " ")
		output = append(output, &numberSet{left: left, right: right})
	}
	return output
}
