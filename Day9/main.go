package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	part1 := sumOfLowPoints("input.txt")
	//part2 := part2("input.txt")
	fmt.Printf("Sum of lowest:  %v \n", part1)
}

func sumOfLowPoints(s string) int {
	result := 0
	elementsSeen := 0
	intArray := readInFileToIntArray(s)
	fmt.Println(len(intArray[0]))
	for i := 0; i < len(intArray); i++ {
		for j := 0; j < len(intArray[0]); j++ {
			isLowest := isLowestPoint(intArray, i, j)
			elementsSeen++
			if isLowest {
				result++
				result += intArray[i][j]
			}

		}
	}
	fmt.Println(elementsSeen)
	return result
}
func isLowestPoint(intArray [][]int, i int, j int) bool {
	val := intArray[i][j]
	//isLowerThanTopLeft := isLower(intArray, i-1, j-1, val)
	isLowerThanTop := isLower(intArray, i, j-1, val)
	//isLowerThanTopRight := isLower(intArray, i+1, j-1, val)
	isLowerThanRight := isLower(intArray, i+1, j, val)
	//isLowerThanBottomRight := isLower(intArray, i+1, j+1, val)
	isLowerThanBottom := isLower(intArray, i, j+1, val)
	//isLowerThanBottomLeft := isLower(intArray, i-1, j+1, val)
	isLowerThanLeft := isLower(intArray, i-1, j, val)

	//return isLowerThanTopLeft && isLowerThanTop && isLowerThanTopRight &&
	//	isLowerThanRight && isLowerThanBottomRight && isLowerThanBottom &&
	//	isLowerThanBottomLeft && isLowerThanLeft
	return isLowerThanTop && isLowerThanBottom && isLowerThanLeft && isLowerThanRight
}

func isLower(intArray [][]int, i, j, val int) bool {
	output := false
	if i < 0 || j < 0 || i >= len(intArray) || j >= len(intArray[0]) {
		output = true
	} else {
		output = val < intArray[i][j]
	}
	return output
}

func readInFileToIntArray(name string) [][]int {
	output := [][]int{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		intArr := []int{}
		for _, char := range line {

			val := int(char - '0')
			intArr = append(intArr, val)
		}

		output = append(output, intArr)
	}
	return output
}
