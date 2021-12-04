package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1 := powerConsumption("input2.txt")
	part2 := lifeSupportRating("input2.txt")
	fmt.Printf("Power Consumption: %v \n", part1)
	fmt.Printf("Life Support: %v \n", part2)
}

func lifeSupportRating(s string) int {
	lines := readInFileToIntArray(s)
	lenOfBinary := len(lines[0])
	oxygenRating := getRating(lines, lenOfBinary, true)
	co2Rating := getRating(lines, lenOfBinary, false)
	return oxygenRating * co2Rating
}

func getRating(lines []string, lenOfBinary int, isOxygen bool) int {
	topChar := byte('1')
	botChar := byte('0')
	if !isOxygen {
		topChar = '0'
		botChar = '1'
	}
	for i := 0; i < lenOfBinary; i++ {
		weight := getWeightAtPosition(i, lines)
		// remove not needed vals
		newLines := []string{}

		for _, item := range lines {
			if weight >= 0 && item[i] == topChar {
				newLines = append(newLines, item)
			}

			if weight < 0 && item[i] == botChar {
				newLines = append(newLines, item)
			}
		}
		//set lines to new lines
		lines = newLines
		// if 1 item is in the list return
		if len(lines) == 1 {
			intVal, _ := strconv.ParseInt(lines[0], 2, 64)
			return int(intVal)
		}
	}
	// err val
	return -1
}

func getWeightAtPosition(pos int, lines []string) int {
	weight := 0
	// get weight of col
	for _, item := range lines {
		if item[pos] == '1' {
			weight++
		} else {
			weight--
		}
	}
	return weight
}

func getBinaryWeightMap(stringArray []string) map[int]int {
	intMap := make(map[int]int)
	lenOfBinary := len(stringArray[0])

	for _, stringVal := range stringArray {
		i := 0
		for i < lenOfBinary {
			if stringVal[i] == '1' {
				intMap[i]++
			} else {
				intMap[i]--
			}
			i++
		}
	}
	return intMap
}

func powerConsumption(s string) int {
	lines := readInFileToIntArray(s)
	intMap := getBinaryWeightMap(lines)
	gamma := 0
	epsilon := 0
	for i := 0; i <= len(intMap)-1; i++ {
		counter := intMap[i]
		if counter > 0 {
			gamma = (gamma << 1) | 1
			epsilon = (epsilon << 1)
		} else {
			gamma = (gamma << 1)
			epsilon = (epsilon << 1) | 1
		}
	}
	output := gamma * epsilon
	//outputBinaryStr := strconv.FormatInt(int64(output), 2)
	//fmt.Println(outputBinaryStr)
	return output
}

func readInFileToIntArray(name string) []string {
	output := []string{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		output = append(output, line)
	}
	return output
}
