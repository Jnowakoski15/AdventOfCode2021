package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type direction struct {
	direction string
	movement  int
}

type cords struct {
	x   int
	y   int
	aim int
}

func main() {
	part1 := finalCordMult("input.txt")
	part2 := finalCordMult2("input.txt")
	fmt.Printf("Cords Multiplied: %v \n", part1)
	fmt.Printf("Part2 Cords Multiplied: %v \n", part2)
}

func finalCordMult2(s string) int {
	cords := cords{x: 0, y: 0, aim: 0}
	dirArray := readInFileToIntArray(s)
	for _, item := range dirArray {

		switch item.direction {
		case "forward":
			cords.x += item.movement
			cords.y += cords.aim * item.movement
		case "down":
			cords.aim += item.movement
		case "up":
			cords.aim -= item.movement
		}
	}

	return cords.x * cords.y
}

func finalCordMult(s string) int {
	cords := cords{x: 0, y: 0}
	dirArray := readInFileToIntArray(s)
	for _, item := range dirArray {

		switch item.direction {
		case "forward":
			cords.x += item.movement
		case "down":
			cords.y += item.movement
		case "up":
			cords.y -= item.movement
		}
	}

	return cords.x * cords.y
}

func readInFileToIntArray(name string) []direction {
	output := []direction{}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		txtVals := strings.Split(line, " ")
		movementInt, _ := strconv.Atoi(txtVals[1])
		dir := direction{direction: txtVals[0], movement: movementInt}
		output = append(output, dir)
	}
	return output
}
