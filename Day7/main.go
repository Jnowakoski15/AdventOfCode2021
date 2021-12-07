package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	part1 := getLeastFuel("input.txt")
	part2 := getLeastFuel2("input.txt")
	fmt.Printf("Least Fuel Used:  %v \n", part1)
	fmt.Printf("Least Fuel Used 2 : %v \n", part2)
}

func getLeastFuel2(s string) int {
	positions, maxPos := readInFileToIntArray(s)
	leastFuelUsed := math.MaxInt64
	for i := 0; i < maxPos; i++ {
		fuelUsed := 0
		for _, val := range positions {
			positionFuelUsed := i - val
			if positionFuelUsed < 0 {
				positionFuelUsed *= -1
			}
			fuelUsed += getNewPositionFuelCost(positionFuelUsed)
		}

		if leastFuelUsed > fuelUsed {
			leastFuelUsed = fuelUsed
		}
	}
	return leastFuelUsed
}

func getNewPositionFuelCost(positionFuelUsed int) int {
	output := 0
	for i := 1; i <= positionFuelUsed; i++ {
		output += i
	}
	return output
}

func getLeastFuel(s string) int {
	positions, maxPos := readInFileToIntArray(s)
	leastFuelUsed := math.MaxInt64
	for i := 0; i < maxPos; i++ {
		fuelUsed := 0
		for _, val := range positions {
			positionFuelUsed := i - val
			if positionFuelUsed < 0 {
				positionFuelUsed *= -1
			}
			fuelUsed += positionFuelUsed
		}

		if leastFuelUsed > fuelUsed {
			leastFuelUsed = fuelUsed
		}
	}
	return leastFuelUsed
}

func readInFileToIntArray(name string) ([]int, int) {
	output := []int{}
	maxInt := 0
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numArry := strings.Split(lines[0], ",")

	for _, numStr := range numArry {
		val, _ := strconv.Atoi(numStr)
		if val > maxInt {
			maxInt = val
		}
		output = append(output, val)
	}
	return output, maxInt
}
